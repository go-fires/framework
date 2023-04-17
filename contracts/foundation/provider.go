package foundation

type Provider interface {
	Register()
	Terminate()
	Boot()
}

// UnimplementedProvider is a default implementation of the Provider interface.
type UnimplementedProvider struct {
}

func (u *UnimplementedProvider) Register() {
}

func (u *UnimplementedProvider) Terminate() {
}

func (u *UnimplementedProvider) Boot() {
}
