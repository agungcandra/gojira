package service

import (
	"regexp"

	"github.com/agungcandra/gojira"
	"github.com/agungcandra/gojira/entity"
)

const (
	ExcludePush            = `(master|release\-|deploy\-)`
	MergeRequestOpenAction = "open"
)

type HookService struct {
	transitionService TransitionInterface
	extractorService  ExtractorInterface
	workflowService   WorkflowService
	credential        *entity.Credential
}

func NewHookService(transitionService TransitionInterface,
	extractorService ExtractorInterface,
	workflowService WorkflowService,
	credential *entity.Credential) *HookService {
	return &HookService{
		transitionService: transitionService,
		extractorService:  extractorService,
		workflowService:   workflowService,
		credential:        credential,
	}
}

func (ms *HookService) MergeRequest(mergeRequest *entity.MergeRequestRequest) []error {
	if mergeRequest.ObjectAttribute.Action != MergeRequestOpenAction {
		return nil
	}

	issues := ms.extractorService.ExtractFromMergeRequest(&mergeRequest.ObjectAttribute)
	return ms.UpdateIssues(issues, gojira.CodeReviewState)
}

func (ms *HookService) Push(pushRequest *entity.PushRequest) []error {
	reg := regexp.MustCompile(ExcludePush)
	if excluded := reg.MatchString(pushRequest.Ref); excluded {
		return nil
	}

	issues := ms.extractorService.ExtractFromPushRequest(pushRequest)
	return ms.UpdateIssues(issues, gojira.InProgressState)
}

func (ms *HookService) UpdateIssues(issues []string, state string) []error {
	ch := make(chan error, len(issues))
	result := make([]error, 0)

	for _, issue := range issues {
		ch <- ms.transitionService.UpdateState(issue, state, ms.credential)
	}

	for range issues {
		err := <-ch
		if err != nil {
			result = append(result, err)
		}
	}

	return result
}
