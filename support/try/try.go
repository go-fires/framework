package try

type Try struct {
	fn     func()
	catchs []func(interface{})
}

func NewTry(fn func()) *Try {
	return &Try{
		fn:     fn,
		catchs: []func(interface{}){},
	}
}

func (t *Try) Catch(fn func(interface{})) *Try {
	t.catchs = append(t.catchs, fn)

	return t
}

func (t *Try) Finally(fns ...func()) {
	defer func() {
		if err := recover(); err != nil {
			for _, catch := range t.catchs {
				catch(err)
			}
		}
	}()

	t.fn()

	if len(fns) > 0 {
		fns[0]()
	}
}

func (t *Try) Do() {
	t.Finally()
}
