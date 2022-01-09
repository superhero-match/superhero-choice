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
	"github.com/superhero-match/superhero-choice/cmd/api/service/mapper"
)

// GetChoice fetches choice(like) from cache if it exists.
// In cache, only like choices are stored, no dislikes.
// The key is in the following form --> superheroID.chosenSuperheroID
func (s *service) GetChoice(key string) (*model.Choice, error) {
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
