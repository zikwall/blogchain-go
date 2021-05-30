package actions

// ToDo: Need to bring all response formats to a single structure and format
type (
	MessageResponse struct {
		Status  uint8  `json:"status"`
		Message string `json:"message"`
	}
	Response struct {
		Response interface{} `json:"response"`
	}
)
