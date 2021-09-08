package service

import (
	"errors"
	"strings"

	"github.com/agungcandra/gojira/entity"
)

type WorkflowService struct {
	workflow entity.Workflow
	request  RequestInterface
}

func NewWorkflowService(workflow entity.Workflow, request RequestInterface) *WorkflowService {
	return &WorkflowService{
		workflow: workflow,
		request:  request,
	}
}

func (w *WorkflowService) Find(issueString, target string, cred *entity.Credential) ([]string, error) {
	issue, err := w.request.GetIssue(issueString, cred)
	if err != nil {
		return nil, err
	}

	issuePrefix := w.getIssuePrefix(issueString)
	workflow := w.workflow[issuePrefix][target]
	if workflow == nil {
		return nil, errors.New("workflow is not supported")
	}

	steps := workflow.Steps[issue.CurrentStatus()]
	return steps, nil
}

func (w *WorkflowService) getIssuePrefix(issue string) string {
	issueSplit := strings.Split(issue, "-")
	return strings.ToUpper(issueSplit[0])
}
