package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Event struct {
	MatchID    string    `json:"match_id"`
	Timestamp  time.Time `json:"timestamp"`
	TeamA      string    `json:"team_a"`
	TeamB      string    `json:"team_b"`
	ScoreTeamA int       `json:"score_team_a"`
	ScoreTeamB int       `json:"score_team_b"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Captures interrupt signals for shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	broker := "localhost:9093"
	topic := "football-example"

	// Group of goroutines to run consumers and producer
	g, gctx := errgroup.WithContext(ctx)

	// Inicia os dois consumers em goroutines separadas
	g.Go(func() error {
		return NotifyConsumer(ctx, logger, broker, topic)
	})

	g.Go(func() error {
		return AnalyticsConsumer(ctx, logger, broker, topic)
	})

	g.Go(func() error {
		scoreA := 1
		scoreB := 1

		// Produces a new event every 5 seconds
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-gctx.Done():
				return gctx.Err() // Exit if context is canceled
			case <-ticker.C:
				event := Event{
					MatchID:    "Team A x Team B",
					Timestamp:  time.Now().UTC(),
					TeamA:      "Time A",
					TeamB:      "Time B",
					ScoreTeamA: scoreA,
					ScoreTeamB: scoreB,
				}
				if err := Producer(ctx, logger, event, broker, topic); err != nil {
					return err
				}
			}
		}
	})

	// Shutdown
	select {
	case <-signalChan: // When receiving an interrupt signal
		logger.Info("Shutdown signal received")
		cancel() // Cancel context to stop goroutines
	case <-gctx.Done(): // If the context is canceled by error
	}

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil && err != context.Canceled {
		logger.Error("Error while running", slog.String("error", err.Error()))
	} else {
		logger.Info("All consumers and producers finished successfully")
	}
}
