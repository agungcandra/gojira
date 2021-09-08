package request

type TransitionRequest struct {
	Transition Transition `json:"transition"`
}

type Transition struct {
	ID string `json:"id"`
}

func NewTransitionRequest(id string) *TransitionRequest {
	transition := Transition{ID: id}
	return &TransitionRequest{
		Transition: transition,
	}
}
