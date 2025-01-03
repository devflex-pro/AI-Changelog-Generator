package git_provider_github

type GitHubCommitResponse struct {
	SHA    string `json:"sha"`
	Commit struct {
		Author struct {
			Name string `json:"name"`
		} `json:"author"`
		Committer struct {
			Name string `json:"name"`
		} `json:"committer"`
		Message string `json:"message"`
	} `json:"commit"`
}
