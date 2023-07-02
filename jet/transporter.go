package jet

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

type Transporter interface {
	Send(data []byte) ([]byte, error)
}

// HttpTransporterConfig is the configuration for the HttpTransporter.
type HttpTransporterConfig struct {
	Url       string
	Timeout   time.Duration
	UserAgent string
}

func (o *HttpTransporterConfig) init() {
	if o.Timeout == 0 {
		o.Timeout = time.Second * 5
	}

	if o.UserAgent == "" {
		o.UserAgent = fmt.Sprintf("Jet/%d Go/%s", 1, runtime.Version())
	}
}

type HttpTransporter struct {
	config *HttpTransporterConfig

	httpClient *http.Client
}

var _ Transporter = (*HttpTransporter)(nil)

func NewHttpTransporter(config *HttpTransporterConfig) *HttpTransporter {
	config.init()

	return &HttpTransporter{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func (t *HttpTransporter) Send(data []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, t.config.Url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", t.config.UserAgent)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
