package fxloader

import (
	"rices/apis/middlewares"
	"rices/apis/routers"
	"rices/common/logger"
	"rices/core/adapters"
	"rices/core/adapters/repository"

	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadEngine()...),
		fx.Options(loadLogger()...),
	}
}

func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(
			adapters.NewPgsql,
		),
		fx.Invoke(func(db *adapters.Pgsql) error {
			return db.Connect()
		}),
		fx.Provide(repository.NewRepositoryUser),
	}
}

func loadUseCase() []fx.Option {
	return []fx.Option{}
}

func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		fx.Provide(middlewares.NewMiddlewareCors),
	}
}

func loadLogger() []fx.Option {
	return []fx.Option{
		fx.Provide(logger.NewLogger),
	}
}
