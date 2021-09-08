package entity

type PushRequest struct {
	Ref     string   `json:"ref"`
	Project Project  `json:"project"`
	Commits []Commit `json:"commits"`
}
