package service

import "github.com/agungcandra/gojira/entity"

type TransitionInterface interface {
	GetTransition(issue string, state string, credential *entity.Credential) (*entity.Transition, error)
	UpdateState(issue string, state string, credential *entity.Credential) error
}

type HookInterface interface {
	MergeRequest(mergeRequest *entity.MergeRequestRequest) error
	Push(pushRequest *entity.PushRequest) []error
}

type ExtractorInterface interface {
	ExtractFromMergeRequest(mr *entity.MergeRequestAttribute) []string
	ExtractFromPushRequest(push *entity.PushRequest) []string
}

type WorkflowInterface interface {
	Find(issueString, target string, cred *entity.Credential) ([]string, error)
}
