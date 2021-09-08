package entity

type Project struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	WebURL        string `json:"web_url"`
	GitSSHUrl     string `json:"git_ssh_url"`
	GitHTTPUrl    string `json:"git_http_url"`
	Namespace     string `json:"namespace"`
	DefaultBranch string `json:"default_branch"`
}
