package service

import (
	"github.com/superhero-choice/cmd/api/model"
	"github.com/superhero-choice/cmd/api/service/mapper"
)

// GetChoice fetches choice(like) from cache if it exist.
// In cache, only like choices are stored, no dislikes.
// The key is in the following form --> superheroID.chosenSuperheroID
func (s *Service) GetChoice(key string) (*model.Choice, error) {
	res, err := s.Cache.GetChoice(key)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	choice := mapper.MapCacheChoiceToResult(*res)

	return &choice, nil
}
