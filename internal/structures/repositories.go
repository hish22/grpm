package structures

type license struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	Url    string `json:"url"`
}

type Repository struct {
	ID                  int      `json:"id"`
	NodeID              string   `json:"node_id"`
	Name                string   `json:"name"`
	FullName            string   `json:"full_name"`
	Private             bool     `json:"private"`
	Owner               Owner    `json:"owner"`
	HtmlUrl             string   `json:"html_url"`
	Description         string   `json:"description"`
	Stars               int      `json:"stargazers_count"`
	Watchers            int      `json:"watchers_count"`
	Forks               int      `json:"forks"`
	ProgrammingLanguage string   `json:"language"`
	License             license  `json:"license"`
	Topics              []string `json:"topics"`
}

type Owner struct {
	ID       int    `json:"id"`
	Username string `json:"login"`
	Url      string `json:"url"`
	HtmlUrl  string `json:"html_url"`
	Type     string `json:"type"`
}

type Repositories struct {
	TotalCount   int          `json:"total_count"`
	Repositories []Repository `json:"items"`
}
