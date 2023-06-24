package recovery

import (
	"github.com/go-fires/fires/contracts/debug/recovery"
)

type Handler struct {
}

var _ recovery.Handler = (*Handler)(nil)

func (p Handler) Report(v interface{}) {
	panic(v)
}
