package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rices/apis/routers"
	"rices/common/configs"
	"rices/fxloader"

	"go.uber.org/fx"
)

func init() {
	var pathConfig string
	flag.StringVar(&pathConfig, "configs", "configs/configs.json", "path config")
	flag.Parse()
	configs.LoadConfig(pathConfig)
}

func main() {
	app := fx.New(
		fx.Provide(configs.Get),
		fx.Options(fxloader.Load()...),
		fx.Invoke(serverLifecycle),
		fx.Options(), // No need for conditional logic with nopLogger
	)

	// Run the application
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err, "Error starting application")
	}

	// Wait for an interrupt signal to gracefully shut down the application
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Shut down the application gracefully
	if err := app.Stop(context.Background()); err != nil {
		log.Fatal(err, "Error stopping application")
	}
}

func serverLifecycle(lc fx.Lifecycle, apiRouter *routers.ApiRouter, cf *configs.Configs) {
	server := &http.Server{
		Addr:    ":" + cf.Port,
		Handler: apiRouter.Engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal(err, "Cannot start server,address")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			//todo user log custom
			//	log.Infof("Stopping backend server.", cf.Port)
			return server.Shutdown(ctx)
		},
	})
}
