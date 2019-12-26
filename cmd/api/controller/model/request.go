package model

// Request holds the data that comes with the /choice request.
type Request struct {
	SuperheroID       string `json:"superheroID"`
	ChosenSuperheroID string `json:"chosenSuperheroID"`
	Choice            int64  `json:"choice"`
}
