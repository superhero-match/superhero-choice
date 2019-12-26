package service

import (
	model2 "github.com/superhero-choice/cmd/api/controller/model"
	mapper2 "github.com/superhero-choice/cmd/api/controller/service/mapper"
)

// StoreChoice publishes new choice on Kafka topic for it to be
// consumed by consumer and stored in Cache.
func (s *Service) StoreChoice(c model2.Choice) error {
	return s.Producer.StoreChoice(mapper2.MapAPIChoiceToProducer(c))
}
