package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

func AnalyticsConsumer(ctx context.Context, logger *slog.Logger, broker, topic string) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	defer r.Close()
	logger.Info("AnalyticsConsumer started. Waiting for events...")

	var totalScore, teamATotalScore, teamBTotalScore int
	// Function to update counters
	updateScores := func(event Event) {
		totalScore += event.ScoreTeamA + event.ScoreTeamB // Total Goals in the Match
		teamATotalScore += event.ScoreTeamA               // Total de gols Time A
		teamBTotalScore += event.ScoreTeamB               // Total de gols Time B

		logger.Info("AnalyticsConsumer:",
			"Match", event.MatchID,
			"Total Goals", totalScore,
			"Total Goals: "+event.TeamA, teamATotalScore,
			"Total Goals: "+event.TeamB, teamBTotalScore)
	}

	for {
		// Verifica se o contexto foi cancelado
		select {
		case <-ctx.Done():
			logger.Info("Context canceled, stopping consumer...")
			return ctx.Err() // Retorna o erro associado ao cancelamento do contexto

		default:
			// Tenta ler uma nova mensagem
			m, err := r.ReadMessage(ctx)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded {
					logger.Info("Context deadline exceeded or canceled, stopping consumer...")
					return err
				}
				logger.Error("could not read message: ", slog.String("error", err.Error()))
				return err
			}

			// Transforma o evento de json para struct em go
			var event Event
			if err := json.Unmarshal(m.Value, &event); err != nil {
				logger.Error("unmarshal", slog.String("error", err.Error()))
				return err
			}

			// Atualiza os contadores ao receber um novo evento
			updateScores(event)
		}
	}
}
