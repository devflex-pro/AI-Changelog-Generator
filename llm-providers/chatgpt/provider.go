package llm_provider_chatgpt

type provider struct{}

func New() *provider {
	return &provider{}
}
