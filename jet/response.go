package jet

import "fmt"

type Error struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s, data: %v", e.Code, e.Message, e.Data)
}

type Response struct {
	ID     string      `json:"id,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Error  *Error      `json:"error,omitempty"`
}
