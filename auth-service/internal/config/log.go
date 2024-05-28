package config

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer log.Sync()

	return log
}
