package oauth2

type Providers struct {
	providers map[string]Provider
}

func (p *Providers) Get(name string) Provider { //nolint:ireturn
	return p.providers[name]
}

func NewProviders(providers map[string]Provider) *Providers {
	return &Providers{
		providers: providers,
	}
}
