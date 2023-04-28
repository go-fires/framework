package cache

import (
	"sync"

	"github.com/go-fires/framework/config"
	"github.com/go-fires/framework/contracts/cache"
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/redis"
)

type Manager struct {
	container container.Container
	config    *Config
	mu        sync.Mutex

	stores map[string]cache.Repository
}

func NewManager(container container.Container) *Manager {
	m := &Manager{
		container: container,
		config:    container.MustGet("config").(*config.Config).Get("cache").(*Config),
		stores:    make(map[string]cache.Repository),
	}

	return m
}

func (m *Manager) Store(names ...string) cache.Repository {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = m.getDefaultStore()
	}

	return m.store(name)
}

func (m *Manager) store(name string) cache.Repository {
	if store, ok := m.stores[name]; ok {
		return store
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.stores[name] = m.resolve(name)

	return m.stores[name]
}

func (m *Manager) resolve(name string) cache.Repository {
	cfg, ok := m.config.Stores[name]
	if !ok {
		panic("cache store not found")
	}

	switch cfg := cfg.(type) {
	case *MemoryStoreConfig:
		return m.createMemoryStore(cfg)
	case *RedisStoreConfig:
		return m.createRedisStore(cfg)
	default:
		panic("cache store not found")
	}
}

func (m *Manager) createMemoryStore(config *MemoryStoreConfig) cache.Repository {
	return m.repository(NewMemoryStore(config))
}

func (m *Manager) createRedisStore(config *RedisStoreConfig) cache.Repository {
	return m.repository(
		NewRedisStore(
			m.container.MustGet("redis").(*redis.Manager).Connect(config.GetConnection()),
			WithRedisStorePrefix(config.GetPrefix()),
			WithRedisStoreSerializable(config.GetSerializer()),
		),
	)
}

func (m *Manager) repository(store cache.Store) cache.Repository {
	return NewRepository(store)
}

func (m *Manager) getDefaultStore() string {
	return m.config.Default
}
