package mapper

import (
	"github.com/superhero-choice/cmd/api/model"
	producer "github.com/superhero-choice/internal/producer/model"
)

// MapAPIChoiceToProducer maps API Choice model to Producer model.
func MapAPIChoiceToProducer(apiChoice model.Choice) producer.Choice {
	return producer.Choice{
		ID:                apiChoice.ID,
		Choice:            apiChoice.Choice,
		SuperheroID:       apiChoice.SuperheroID,
		ChosenSuperheroID: apiChoice.ChosenSuperheroID,
		CreatedAt:         apiChoice.CreatedAt,
	}
}
