package common

type Response struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func SuccessResponse() *Response {
	return &Response{
		"success",
		"",
	}
}

func ErrorResponse(e error) *Response {
	return &Response{
		"error",
		e.Error(),
	}
}

func ErrorResponseStr(errorMessage string) *Response {
	return &Response{
		"error",
		errorMessage,
	}
}
