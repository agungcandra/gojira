package entity

type Issue struct {
	ID     string     `json:"id"`
	Key    string     `json:"key"`
	Fields IssueField `json:"fields"`
}

type IssueField struct {
	Status IssueStatus `json:"status"`
}

type IssueStatus struct {
	Name string `json:"name"`
}

func (i *Issue) CurrentStatus() string {
	return i.Fields.Status.Name
}
