package logger

import (
	"github.com/fkihai/payflow/internal/infrastructure/config"
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(cfg *config.AppConfig) error {
	var err error
	if cfg.Env == "development" {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}

	Log.Info("logger initialized",
		zap.String("env", cfg.Env),
	)
	return err
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
