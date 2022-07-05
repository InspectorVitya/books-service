package main

import (
	"books-service/internal/app"
	"books-service/internal/config"
	"books-service/internal/server"
	"books-service/internal/server/grpcserver"
	"books-service/internal/storage/mysql"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	out := zerolog.NewConsoleWriter()
	out.Out = os.Stderr
	logger := zerolog.New(out)

	cfg, err := config.New()
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	db, err := mysql.New(cfg.HostDB, cfg.NameDB, cfg.UserDB, cfg.PasswordDB, cfg.PortDB)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	bookApp := app.New(db)

	svrGRPC := grpcserver.New(bookApp, cfg.GRPSPort)
	svr := server.New(svrGRPC)

	go func() {
		logger.Info().Msg("grpc server start...")
		if err := svr.GRPC.Start(); err != nil {
			logger.Fatal().Err(err).Send()
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-stop
	svr.GRPC.Stop()
	err = db.Stop()
	if err != nil {
		logger.Err(err).Send()
	}
	logger.Info().Msg("service stopped...")
}
