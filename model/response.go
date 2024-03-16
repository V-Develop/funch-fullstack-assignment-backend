package model

type ResponseWithoutPayload struct {
	Status ResponseHeader `json:"status"`
}

type ResponseWithPayload struct {
	Status  ResponseHeader `json:"status"`
	Payload interface{}    `json:"payload"`
}

type ResponseHeader struct {
	HttpStatus int    `json:"http_status"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}
