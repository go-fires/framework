package rediser

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

type Manager struct {
	config   *Config
	resolved map[string]redis.Cmdable
	mu       sync.Mutex
}

type Opts func(*Manager) *Manager

// New a redis manager.
func New(config *Config, opts ...Opts) *Manager {
	m := &Manager{
		config:   config,
		resolved: make(map[string]redis.Cmdable, len(config.Connections)),
	}

	for _, opt := range opts {
		m = opt(m)
	}

	return m
}

// Connect get a redis client.
func (m *Manager) Connect(name ...string) redis.Cmdable {
	if len(name) > 0 {
		return m.resolve(name[0])
	}

	return m.resolve(m.config.Default)
}

// resolve client.
func (m *Manager) resolve(name string) redis.Cmdable {
	if rdb, ok := m.resolved[name]; ok {
		return rdb
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.config.Connections[name]; !ok {
		panic("connection " + name + " is not defined")
	}

	var rdb redis.Cmdable
	switch m.config.Connections[name].(type) {
	case *redis.Options:
		rdb = redis.NewClient(m.config.Connections[name].(*redis.Options))
		break
	case *redis.ClusterOptions:
		rdb = redis.NewClusterClient(m.config.Connections[name].(*redis.ClusterOptions))
		break
	case *redis.RingOptions:
		rdb = redis.NewRing(m.config.Connections[name].(*redis.RingOptions))
		break
	default:
		panic("connection " + name + " is not defined")
	}

	m.resolved[name] = rdb

	return rdb
}
