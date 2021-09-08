package entity

type Commit struct {
	ID        string       `json:"id"`
	Message   string       `json:"message"`
	Title     string       `json:"title"`
	Timestamp string       `json:"timestamp"`
	Author    CommitAuthor `json:"author"`
	Added     []string     `json:"added"`
	Modified  []string     `json:"modified"`
	Removed   []string     `json:"removed"`
}

type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
