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
package cache

import (
	"github.com/go-redis/redis"

	"github.com/superhero-match/superhero-choice/internal/cache/model"
)

// GetChoice fetches choice(like) from cache if it exists.
// In cache, only like choices are stored, no dislikes.
// The key is in the following form --> superheroID.chosenSuperheroID
func (c *cache) GetChoice(key string) (*model.Choice, error) {
	var choice model.Choice

	res, err := c.Redis.Get(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	if err := choice.UnmarshalBinary([]byte(res)); err != nil {
		return nil, err
	}

	return &choice, nil
}
