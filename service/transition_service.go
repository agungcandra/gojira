package service

import (
	"errors"

	"github.com/agungcandra/gojira/entity"
)

type TransitionService struct {
	request  RequestInterface
	workflow WorkflowInterface
}

func NewTransitionService(request RequestInterface,
	workflow WorkflowInterface) *TransitionService {
	return &TransitionService{
		request:  request,
		workflow: workflow,
	}
}

func (t *TransitionService) GetTransition(issue string, state string, credential *entity.Credential) (*entity.Transition, error) {
	transitions, err := t.request.GetAvailableTransition(issue, credential)
	if err != nil {
		return nil, err
	}

	for _, transition := range transitions {
		if state == transition.Name && transition.IsAvailable {
			return &transition, nil
		}
	}

	return nil, errors.New("not available")
}

func (t *TransitionService) UpdateState(issue string, state string, credential *entity.Credential) error {
	steps, err := t.workflow.Find(issue, state, credential)
	if err != nil {
		return err
	}

	for _, state := range steps {
		transition, err := t.GetTransition(issue, state, credential)
		if err != nil {
			return err
		}

		t.request.UpdateJiraState(issue, transition, credential)
	}

	return nil
}
