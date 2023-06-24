package recovery

type PanicHandler struct{}

var _ Handler = (*PanicHandler)(nil)

func (p *PanicHandler) Report(v interface{}) {
	if p.ShouldReport(v) {
		panic(v)
	}
}

func (p *PanicHandler) ShouldReport(v interface{}) bool {
	return true
}
