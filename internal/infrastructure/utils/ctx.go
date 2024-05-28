package utils

import (
	"context"
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/contants"
)

func GetConfig(ctx context.Context) *config.Config {
	contextConfig := ctx.Value(contants.AppConfigName)
	return contextConfig.(*config.Config)
}
