package cache

import (
	"github.com/go-fires/framework/contracts/cache"
	"github.com/go-fires/framework/facade"
	"sync"
)

type Config struct {
	Default string
	Stores  map[string]cache.StoreConfigable
}

type Manager struct {
	config *Config
	mu     sync.Mutex

	stores map[string]cache.Repository
}

func NewManager(config *Config) *Manager {
	m := &Manager{
		config: config,
		stores: make(map[string]cache.Repository),
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
	config, ok := m.config.Stores[name]
	if !ok {
		panic("cache store not found")
	}

	switch config.(type) {
	case *MemoryStoreConfig:
		return m.createMemoryStore(config.(*MemoryStoreConfig))
	case *RedisStoreConfig:
		return m.createRedisStore(config.(*RedisStoreConfig))
	default:
		panic("cache store not found")
	}
}

func (m *Manager) createMemoryStore(config *MemoryStoreConfig) cache.Repository {
	return m.repository(NewMemoryStore())
}

func (m *Manager) createRedisStore(config *RedisStoreConfig) cache.Repository {
	return m.repository(
		NewRedisStore(
			facade.Redis().Connect(config.GetConnection()),
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
