package git_provider_github

type provider struct{}

func New() *provider {
	return &provider{}
}
