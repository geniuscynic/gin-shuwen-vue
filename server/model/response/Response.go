package response

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func CreateSuccessResponse(data interface{}) *Response {
	return &Response{
		Status:  200,
		Message: "",
		Success: true,
		Data:    data,
	}

}

func CreateFailedResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Success: false,
	}

}
