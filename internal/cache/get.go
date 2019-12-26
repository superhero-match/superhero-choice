package cache

import (
	"github.com/superhero-choice/internal/cache/model"
)

// GetChoice fetches choice(like) from cache if it exist.
// In cache, only like choices are stored, no dislikes.
// The key is in the following form --> superheroID.chosenSuperheroID
func (c *Cache) GetChoice(key string) (model.Choice, error) {
	var choice model.Choice

	res, err := c.Redis.Get(key).Result()
	if err != nil {
		return choice, err
	}

	if err := choice.UnmarshalBinary([]byte(res)); err != nil {
		return choice, err
	}

	return choice, nil
}
