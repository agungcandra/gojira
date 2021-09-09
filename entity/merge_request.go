package entity

type MergeRequestRequest struct {
	ObjectKind      string                `json:"object_kind"`
	User            User                  `json:"user"`
	Project         Project               `json:"project"`
	ObjectAttribute MergeRequestAttribute `json:"object_attributes"`
}

type MergeRequestAttribute struct {
	ID              int    `json:"id"`
	TargetBranch    string `json:"target_branch"`
	SourceBranch    string `json:"source_branch"`
	SourceProjectID int    `json:"source_project_id"`
	Title           string `json:"title"`
	LastCommit      Commit `json:"last_commit"`
	Action          string `json:"action"`
}
