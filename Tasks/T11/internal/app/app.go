package app

import (
	"WBL2/Tasks/T11/internal/adapters/cache"
	httpadapter "WBL2/Tasks/T11/internal/adapters/http"
	"WBL2/Tasks/T11/internal/domain/usecase"
	"context"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	logger *zap.Logger
	server httpadapter.AdapterHTTP
)

func Start(ctx context.Context) {
	logger, _ = zap.NewProduction()
	storage := cache.New()
	eventsService := usecase.New(storage)
	server = httpadapter.New(eventsService, logger.Sugar())

	var g errgroup.Group
	g.Go(func() error {
		return server.Start(ctx)
	})

	logger.Sugar().Info("application has started")

	err := g.Wait()
	if err != nil {
		logger.Sugar().Fatalw("http server start failed", zap.Error(err))
	}
}

func Stop(ctx context.Context) {
	if err := server.Stop(ctx); err != nil {
		logger.Sugar().Errorw("http server shutting down error", zap.Error(err))
	}
	logger.Sugar().Info("application has stopped")
}
