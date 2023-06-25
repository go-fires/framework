package logging

import (
	logger2 "github.com/go-fires/fires/x/contracts/logger"
	"sync"
)

type Manager struct {
	config   *Config
	resolved map[string]logger2.Logger
	mu       sync.Mutex

	logger2.Loggerable
}

func NewManager(config *Config) *Manager {
	m := &Manager{
		config:   config,
		resolved: make(map[string]logger2.Logger),
	}

	m.Loggerable = func(level logger2.Level, s string) {
		m.Channel().Log(level, s)
	}

	return m
}

func (m *Manager) Channel(names ...string) logger2.Logger {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = m.getDefaultName()
	}

	return m.Get(name)
}

func (m *Manager) Get(name string) logger2.Logger {
	if logging, ok := m.resolved[name]; ok {
		return logging
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	logging := m.resolve(name)
	m.resolved[name] = logging

	return logging
}

func (m *Manager) resolve(name string) logger2.Logger {
	if log, ok := m.config.Channels[name]; ok {
		return log
	}

	panic("log channel " + name + " is not defined")
}

func (m *Manager) getDefaultName() string {
	return m.config.Default
}
