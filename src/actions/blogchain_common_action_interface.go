package actions

// ToDo: Need to bring all response formats to a single structure and format
type (
	BlogchainCommonResponseAttributes struct {
		Status uint8 `json:"status"`
	}
	BlogchainMessageResponse struct {
		BlogchainCommonResponseAttributes
		Message string `json:"message"`
	}
	BlogchainResponse struct {
		BlogchainCommonResponseAttributes
		Data interface{} `json:"data"`
	}
)
