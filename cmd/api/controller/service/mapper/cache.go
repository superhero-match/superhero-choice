package mapper

import (
	model2 "github.com/superhero-choice/cmd/api/controller/model"
	cache "github.com/superhero-choice/internal/cache/model"
)

// MapCacheChoiceToResult maps cache Choice model to API models that are returned to the user.
func MapCacheChoiceToResult(cachedChoice cache.Choice) model2.Choice {
	return model2.Choice{
		ID:                cachedChoice.ID,
		Choice:            cachedChoice.Choice,
		SuperheroID:       cachedChoice.SuperheroID,
		ChosenSuperheroID: cachedChoice.ChosenSuperheroID,
		CreatedAt:         cachedChoice.CreatedAt,
	}
}
