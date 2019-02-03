package models

type Response struct {
	err errorResponse `json:"error,omitempty"`
}

type errorResponse struct {
	code    int    `json:"code,omitempty"`
	message string `json:"string,omitempty"`
}

func New(code int, message string) Response {
	return Response{
		errorResponse{
			code:    code,
			message: message,
		},
	}
}
