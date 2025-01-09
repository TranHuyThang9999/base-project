package fxloader

import (
	"rices/apis/controllers"
	"rices/apis/middlewares"
	"rices/apis/resources"
	"rices/apis/routers"
	"rices/common/logger"
	"rices/core/adapters"
	"rices/core/adapters/repository"
	"rices/core/services"

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
	return []fx.Option{
		fx.Provide(services.NewJwtService),
		fx.Provide(services.NewUserService),
	}
}

func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		fx.Provide(middlewares.NewMiddlewareCors),
		fx.Provide(controllers.NewUserController),
		fx.Provide(controllers.NewBaseController),
		fx.Provide(resources.NewResource),
		fx.Provide(middlewares.NewMiddlewareJwt),
	}
}

func loadLogger() []fx.Option {
	return []fx.Option{
		fx.Provide(logger.NewLogger),
	}
}
