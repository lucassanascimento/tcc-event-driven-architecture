package main

import (
	"context"
	"log/slog"
	"os"
	"time"
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
	ctx := context.Background()

	broker := "localhost:9093"
	topic := "football-example"
	scoreA := 1
	scoreB := 1

	go NotifyConsumer(ctx, logger, broker, topic)
	go AnalyticsConsumer(ctx, logger, broker, topic)

	for {
	// sends a new event every 5 seconds
		event := Event{
			MatchID:    "Team A x Team B",
			Timestamp:  time.Now().UTC(),
			TeamA:      "Time A",
			TeamB:      "Time B",
			ScoreTeamA: scoreA,
			ScoreTeamB: scoreB,
		}
		Producer(ctx, logger, event, broker, topic)
		time.Sleep(5 * time.Second)
	}
}
