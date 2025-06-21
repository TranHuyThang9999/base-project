package main

import (
	"context"
	"demo_time_sheet_server/apis/routers"
	"demo_time_sheet_server/common/configs"
	"demo_time_sheet_server/fxloader"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go.uber.org/fx"
)

func init() {
	//	utils.GenerateConfigFile() //use dev
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
		fx.Options(),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err, "Error starting application")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

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
			log.Println("Stopping backend server.", cf.Port)
			return server.Shutdown(ctx)
		},
	})
}
