package foundation

import "sync"

var instance *Application
var syncOnce sync.Once

func SetInstance(app *Application) {
	instance = app
}

func GetInstance() *Application {
	if instance == nil {
		syncOnce.Do(func() {
			instance = NewApplication()
		})
	}

	return instance
}
