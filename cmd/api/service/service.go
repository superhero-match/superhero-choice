/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package service

import (
	"github.com/superhero-match/superhero-choice/cmd/api/model"
	"github.com/superhero-match/superhero-choice/internal/cache"
	"github.com/superhero-match/superhero-choice/internal/config"
	"github.com/superhero-match/superhero-choice/internal/producer"
)

// Service interface defines service methods.
type Service interface {
	GetChoice(key string) (*model.Choice, error)
	StoreChoice(c model.Choice) error
}

// service holds all the different services that are used when handling request.
type service struct {
	Producer producer.Producer
	Cache    cache.Cache
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (Service, error) {
	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	return &service{
		Producer: producer.NewProducer(cfg),
		Cache:    c,
	}, nil
}
