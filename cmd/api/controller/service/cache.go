package service

import (
	model2 "github.com/superhero-choice/cmd/api/controller/model"
	mapper2 "github.com/superhero-choice/cmd/api/controller/service/mapper"
)

// GetChoice fetches choice(like) from cache if it exist.
// In cache, only like choices are stored, no dislikes.
// The key is in the following form --> superheroID.chosenSuperheroID
func (s *Service) GetChoice(key string) (*model2.Choice, error) {
	res, err := s.Cache.GetChoice(key)
	if err != nil {
		return nil, err
	}

	choice := mapper2.MapCacheChoiceToResult(res)

	return &choice, nil
}
