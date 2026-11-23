package response

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponseWithMeta(message string, data interface{}, page, limit, total int) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta: &Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	}
}

func NewErrorResponse(message string) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
	}
}
