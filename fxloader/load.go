package fxloader

import (
	"demo_time_sheet_server/apis/controllers"
	"demo_time_sheet_server/apis/middlewares"
	"demo_time_sheet_server/apis/resources"
	"demo_time_sheet_server/apis/routers"
	"demo_time_sheet_server/common/logger"
	"demo_time_sheet_server/core/adapters"
	"demo_time_sheet_server/core/adapters/repository"
	"demo_time_sheet_server/core/services"

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
		fx.Provide(
			adapters.NewRedis,
		),
		fx.Invoke(func(db *adapters.Pgsql) error {
			return db.Connect()
		}),
		fx.Invoke(func(db *adapters.Redis) error {
			return db.Connect()
		}),
		fx.Provide(repository.NewRepositoryUser),
		fx.Provide(repository.NewRepositoryCache),
		fx.Provide(repository.NewRepositoryTransaction),
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
