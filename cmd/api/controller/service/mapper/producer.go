package mapper

import (
	model2 "github.com/superhero-choice/cmd/api/controller/model"
	producer "github.com/superhero-choice/internal/producer/model"
)

// MapAPIChoiceToProducer maps cache Choice model to API models that are returned to the user.
func MapAPIChoiceToProducer(apiChoice model2.Choice) producer.Choice {
	return producer.Choice{
		ID:                apiChoice.ID,
		Choice:            apiChoice.Choice,
		SuperheroID:       apiChoice.SuperheroID,
		ChosenSuperheroID: apiChoice.ChosenSuperheroID,
		CreatedAt:         apiChoice.CreatedAt,
	}
}
