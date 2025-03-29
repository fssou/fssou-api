package routes

import "encoding/json"

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e *ErrorResponse) ToBytes() []byte {
	value, err := json.Marshal(e)
	if err != nil {
		return []byte(`{"code": "500", "message": "Internal Server Error"}`)
	}
	return value
}

func (e *ErrorResponse) ToString() string {
	value := e.ToBytes()
	return string(value)
}
