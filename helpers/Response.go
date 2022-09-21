package helpers

type Response struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}

type AppError struct {
	Code        int         `json:"code,omitempty"`
	MessageData interface{} `json:"messageData"`
}

// ResponseSuccess
func ResponseSuccess(data interface{}) Response {
	res := Response{
		Status: true,
		Data:   data,
		Error:  nil,
	}
	return res
}

// ResponseError
func ResponseError(messageData interface{}, code int) Response {
	error := AppError{
		Code:        code,
		MessageData: messageData,
	}
	res := Response{
		Status: false,
		Data:   nil,
		Error:  error,
	}
	return res
}
