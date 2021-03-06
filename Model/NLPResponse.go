package Model

// NLPResponse is the struct for the response
type NLPResponse struct {
	Intent          string            `json:"intent"`
	Confidence      float32           `json:"confidence"`
	ResponseMessage string            `json:"fulfillmentText"`
	Entities        map[string]string `json:"entities"`
}
