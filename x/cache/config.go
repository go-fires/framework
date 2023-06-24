package cache

// StoreConfigable is the interface that all cache store drivers must implement.
type StoreConfigable interface {
}

type Config struct {
	Default string
	Stores  map[string]StoreConfigable
}
