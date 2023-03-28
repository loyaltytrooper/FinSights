package helpers

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Success bool        `json:"success"`
}

func CreateResponse(data interface{}, err error) Response {
	var response Response
	response.Data = data
	if err != nil {
		response.Error = err.Error()
		response.Success = false
	} else {
		response.Error = nil
		response.Success = true
	}
	return response
}
