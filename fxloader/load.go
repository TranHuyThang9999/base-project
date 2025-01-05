package fxloader

import (
	"rices/apis/routers"
	"rices/core/adapters"

	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadEngine()...),
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
	}
}

func loadUseCase() []fx.Option {
	return []fx.Option{}
}

func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
	}
}
