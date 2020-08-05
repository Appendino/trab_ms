package issue

type Issue struct {
	IssueID     int    `json:"issueId,omitempty"`
	Username    string `json:"username",omitempty"`
	Description string `json:"desc,omitempty"`
}
