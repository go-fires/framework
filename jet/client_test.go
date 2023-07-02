package jet

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Reply string `json:"reply"`
}

type testService struct {
	*Client
}

func newTestService() *testService {
	return &testService{
		Client: New("test", WithTransporter(
			NewHttpTransporter(&HttpTransporterConfig{
				Url: "http://127.0.0.1:4523/m1/345347-0-default/",
			}),
		)),
	}
}

func (c *testService) Hello(req helloRequest) (resp helloResponse, err error) {
	err = c.Invoke("hello", req, &resp)

	return
}

func TestClient_Call(t *testing.T) {
	s := newTestService()

	spew.Dump(s.Hello(helloRequest{
		Name: "aaa",
	}))
}
