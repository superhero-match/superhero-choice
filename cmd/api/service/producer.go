package service

import (
	"github.com/superhero-choice/cmd/api/model"
	"github.com/superhero-choice/cmd/api/service/mapper"
)

// StoreChoice publishes new choice on Kafka topic for it to be
// consumed by consumer and stored in Cache.
func (s *Service) StoreChoice(c model.Choice) error {
	return s.Producer.StoreChoice(mapper.MapAPIChoiceToProducer(c))
}
