package entity

type Transition struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsAvailable bool   `json:"isAvailable"`
}

type TransitionResponse struct {
	Expand      string       `json:"expand"`
	Transitions []Transition `json:"transitions"`
}
