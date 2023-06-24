package recovery

type panicHandler struct{}

var DefaultHandler Handler = &panicHandler{}

func (p *panicHandler) Report(v interface{}) {
	if p.ShouldReport(v) {
		panic(v)
	}
}

func (p *panicHandler) ShouldReport(v interface{}) bool {
	return true
}
