package jet

type Request struct {
	Path   string      `json:"path"`
	Params interface{} `json:"params"`
	ID     string      `json:"id"`
}
