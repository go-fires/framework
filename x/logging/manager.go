package logging

import (
	"sync"

	"github.com/go-fires/fires/contracts/logger"
)

type Manager struct {
	config   *Config
	resolved map[string]logger.Logger
	mu       sync.Mutex

	logger.Loggerable
}

func NewManager(config *Config) *Manager {
	m := &Manager{
		config:   config,
		resolved: make(map[string]logger.Logger),
	}

	m.Loggerable = func(level logger.Level, s string) {
		m.Channel().Log(level, s)
	}

	return m
}

func (m *Manager) Channel(names ...string) logger.Logger {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = m.getDefaultName()
	}

	return m.Get(name)
}

func (m *Manager) Get(name string) logger.Logger {
	if logging, ok := m.resolved[name]; ok {
		return logging
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	logging := m.resolve(name)
	m.resolved[name] = logging

	return logging
}

func (m *Manager) resolve(name string) logger.Logger {
	if log, ok := m.config.Channels[name]; ok {
		return log
	}

	panic("log channel " + name + " is not defined")
}

func (m *Manager) getDefaultName() string {
	return m.config.Default
}
