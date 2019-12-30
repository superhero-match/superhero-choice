package mapper

import (
	"github.com/superhero-choice/cmd/api/model"
	cache "github.com/superhero-choice/internal/cache/model"
)

// MapCacheChoiceToResult maps cache Choice model to API model(external model) that is returned to the user.
func MapCacheChoiceToResult(cachedChoice cache.Choice) model.Choice {
	return model.Choice{
		ID:                cachedChoice.ID,
		Choice:            cachedChoice.Choice,
		SuperheroID:       cachedChoice.SuperheroID,
		ChosenSuperheroID: cachedChoice.ChosenSuperheroID,
		CreatedAt:         cachedChoice.CreatedAt,
	}
}
