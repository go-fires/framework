package jet

type DataFormatter interface {
	EncodeRequest(request Request) interface{}
	EncodeResponse(response Response) interface{}
	DecodeResponse(data []byte, unserialize func([]byte, interface{}) error) (Response, error)
}

const jsonRPCVersion = "2.0"

var DefaultDataFormatter DataFormatter = &JsonRPCDataFormatter{}

type JsonRPCDataFormatter struct{}

type JsonRPCRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      string      `json:"id,omitempty"`
}

type JsonRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type JsonRPCResponse struct {
	JsonRPC string        `json:"jsonrpc"`
	ID      string        `json:"id,omitempty"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JsonRPCError `json:"error"`
}

func (j *JsonRPCDataFormatter) EncodeRequest(request Request) interface{} {
	return JsonRPCRequest{
		JsonRPC: jsonRPCVersion,
		Method:  request.Path,
		Params:  request.Params,
		ID:      request.ID,
	}
}

func (j *JsonRPCDataFormatter) EncodeResponse(response Response) interface{} {
	resp := JsonRPCResponse{
		JsonRPC: jsonRPCVersion,
		ID:      response.ID,
		Result:  response.Result,
	}

	if response.Error != nil {
		resp.Error = &JsonRPCError{
			Code:    response.Error.Code,
			Message: response.Error.Message,
			Data:    response.Error.Data,
		}
	}

	return resp
}

func (j *JsonRPCDataFormatter) DecodeResponse(data []byte, unserialize func([]byte, interface{}) error) (Response, error) {
	var response JsonRPCResponse
	if err := unserialize(data, &response); err != nil {
		return Response{}, err
	}

	resp := Response{
		ID:     response.ID,
		Result: response.Result,
	}

	if response.Error != nil {
		resp.Error = &Error{
			Code:    response.Error.Code,
			Message: response.Error.Message,
			Data:    response.Error.Data,
		}
	}

	return resp, nil
}
