package entity

type WorkflowState string

const (
	WorkflowInProgress WorkflowState = "InProgress"
	WorkflowCodeReview WorkflowState = "CodeReview"
	WorkflowDone       WorkflowState = "Done"
)

type Workflow map[string]map[string]*WorkflowDetail

type WorkflowDetail struct {
	Target string              `yaml:"Target"`
	Steps  map[string][]string `yaml:"Steps"`
}
