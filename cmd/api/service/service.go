package service

import (
	"github.com/superhero-choice/internal/cache"
	"github.com/superhero-choice/internal/config"
	"github.com/superhero-choice/internal/producer"
	"go.uber.org/zap"
)

// Service holds all the different services that are used when handling request.
type Service struct {
	Producer   *producer.Producer
	Cache      *cache.Cache
	Logger     *zap.Logger
	TimeFormat string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (*Service, error) {
	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Service{
		Producer:   producer.NewProducer(cfg),
		Cache:      c,
		Logger:     logger,
		TimeFormat: cfg.App.TimeFormat,
	}, nil
}
