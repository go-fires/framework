package context

import "sync"

var contexts = make(map[string]interface{})
var mu sync.Mutex

func Set(key string, value interface{}) {
	mu.Lock()
	defer mu.Unlock()

	contexts[key] = value
}

func Get(key string) interface{} {
	if context, ok := contexts[key]; ok {
		return context
	}

	return nil
}

func Has(key string) bool {
	_, ok := contexts[key]

	return ok
}

func Delete(key string) {
	mu.Lock()
	defer mu.Unlock()

	delete(contexts, key)
}

func Clear() {
	mu.Lock()
	defer mu.Unlock()

	contexts = make(map[string]interface{})
}
