package foundation

import (
	"fmt"
	"github.com/go-fires/framework/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testProvider struct {
	*container.Container
}

type testService struct {
	Name string
}

func (t *testProvider) Register() {
	t.Singleton("test", func(c *container.Container) interface{} {
		return &testService{
			Name: "test name",
		}
	})
}

func (t *testProvider) Boot() {
	fmt.Println("test provider booted")
}

func (t *testProvider) Terminate() {
	fmt.Println("test provider terminated")
}

func TestApplication_Register(t *testing.T) {
	app := NewApplication()

	app.Register(&testProvider{
		Container: app.Container,
	})

	var test *testService
	assert.Nil(t, app.Get("test", &test))
	assert.Equal(t, "test name", test.Name)

	app.Boot()
	app.Terminate()
}