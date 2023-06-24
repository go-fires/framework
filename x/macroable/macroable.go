package macroable

import (
	"sync"

	"github.com/go-fires/fires/support/helper"
)

type Macroable struct {
	ctx    interface{}
	macros map[string]interface{}
	mu     sync.Mutex
}

func NewMacroable() *Macroable {
	return &Macroable{
		macros: make(map[string]interface{}),
	}
}

func (m *Macroable) Macro(name string, callback interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.macros[name] = callback
}

func (m *Macroable) HasMacro(name string) bool {
	_, ok := m.macros[name]
	return ok
}

func (m *Macroable) WithCtx(ctx interface{}) *Macroable {
	m.ctx = ctx

	return m
}

func (m *Macroable) Call(name string, args ...interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()

	if callback, ok := m.macros[name]; ok {
		return helper.CallWithCtx(m.ctx, callback, args...)
	}

	panic("macro " + name + " is not defined")
}
