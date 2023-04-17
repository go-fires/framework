package hashing

import (
	"fmt"
	"github.com/go-fires/framework/contracts/hashing"
	"sync"
)

type Manager struct {
	config  *Config
	drivers map[string]hashing.Hasher
	rw      *sync.RWMutex
}

type Config struct {
	Driver string
}

// NewManager creates a new hashing manager instance.
// config example:
//
//	config := &Config{
//		Driver: "bcrypt",
//	}
func NewManager(config *Config) *Manager {
	return &Manager{
		config:  config,
		drivers: make(map[string]hashing.Hasher, 5),
		rw:      &sync.RWMutex{},
	}
}

// Driver gets the hasher instance by driver name.
func (m *Manager) Driver(driver ...string) hashing.Hasher {
	if len(driver) > 0 {
		return m.resolve(driver[0])
	}

	return m.resolve(m.getDefaultDriver())
}

// resolve gets the hasher instance by name.
func (m *Manager) resolve(driver string) hashing.Hasher {
	m.rw.Lock()
	defer m.rw.Unlock()

	hasher, ok := m.drivers[driver]
	if ok {
		return hasher
	}

	switch driver {
	case "bcrypt":
		hasher = m.createBcryptHasher()
	case "md5":
		hasher = m.createMd5Hasher()
	case "sha1":
		hasher = m.createSha1Hasher()
	default:
		panic(fmt.Sprintf("hashing driver %s is not supported", driver))
	}

	m.drivers[driver] = hasher

	return hasher
}

// createBcryptHasher creates a new bcrypt hasher instance.
func (m *Manager) createBcryptHasher() hashing.Hasher {
	return NewBcryptHasher()
}

// createMd5Hasher creates a new md5 hasher instance.
func (m *Manager) createMd5Hasher() hashing.Hasher {
	return NewMd5Hasher()
}

// createSha1Hasher creates a new sha1 hasher instance.
func (m *Manager) createSha1Hasher() hashing.Hasher {
	return NewSha1Hasher()
}

// getDefaultDriver gets the default driver name.
func (m *Manager) getDefaultDriver() string {
	return m.config.Driver
}

// Make creates a hash value for the given value.
func (m *Manager) Make(value string) (string, error) {
	return m.Driver().Make(value)
}

// MustMake creates a hash value for the given value.
func (m *Manager) MustMake(value string) string {
	return m.Driver().MustMake(value)
}

// Check checks the given value is equal to the hashed value.
func (m *Manager) Check(value, hashedValue string) bool {
	return m.Driver().Check(value, hashedValue)
}
