/*
  Copyright (C) 2019 - 2021 MWSOFT
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
package mapper

import (
	"github.com/superhero-match/superhero-choice/cmd/api/model"
	cache "github.com/superhero-match/superhero-choice/internal/cache/model"
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
