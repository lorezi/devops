package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger = logger.Named("my-app")
	logger.Info("failed to fetch URL",
		zap.String("url", "https:github.now"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
