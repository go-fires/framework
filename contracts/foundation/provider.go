package foundation

type Provider interface {
	Register()
	Terminate()
	Boot()
}
