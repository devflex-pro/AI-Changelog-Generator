package destination_provider_slack

type provider struct{}

func New() *provider {
	return &provider{}
}
