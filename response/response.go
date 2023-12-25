package response

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Format(message string, data interface{}) response {
	return response{
		Message: message,
		Data:    data,
	}
}
