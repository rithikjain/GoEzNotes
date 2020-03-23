package views

type PingResponse struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}
