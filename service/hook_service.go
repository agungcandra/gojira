package service

import (
	"regexp"

	"github.com/agungcandra/gojira"
	"github.com/agungcandra/gojira/entity"
)

const ExcludePush = `(master|release\-|deploy\-)`

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

func (ms *HookService) MergeRequest(mergeRequest *entity.MergeRequestRequest) error {
	issues := ms.extractorService.ExtractFromMergeRequest(&mergeRequest.ObjectAttribute)

	for _, issue := range issues {
		err := ms.transitionService.UpdateState(issue, gojira.CodeReviewState, ms.credential)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ms *HookService) Push(pushRequest *entity.PushRequest) []error {
	reg := regexp.MustCompile(ExcludePush)
	if excluded := reg.MatchString(pushRequest.Ref); excluded {
		return nil
	}

	issues := ms.extractorService.ExtractFromPushRequest(pushRequest)
	ch := make(chan error, len(issues))
	result := make([]error, 0)

	for _, issue := range issues {
		ch <- ms.transitionService.UpdateState(issue, gojira.InProgressState, ms.credential)
	}

	for range issues {
		err := <-ch
		if err != nil {
			result = append(result, err)
		}
	}

	return result
}
