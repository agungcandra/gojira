package service

import "github.com/agungcandra/gojira/entity"

type RequestInterface interface {
	GetAvailableTransition(issue string, credential *entity.Credential) ([]entity.Transition, error)
	UpdateJiraState(issue string, transition *entity.Transition, credential *entity.Credential) error
	GetIssue(issue string, credential *entity.Credential) (*entity.Issue, error)
}
