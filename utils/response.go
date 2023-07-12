package utils

type ResponseError struct {
	Code    int    `json:"code,omitempty" binding:"omitempty"`
	Message string `json:"message,omitempty" binding:"omitempty"`
	Data    struct {
		Message string `json:"message,omitempty" binding:"omitempty"`
	} `json:"data,omitempty" binding:"omitempty"`
}
type ResponseSuccess struct {
	Code    int         `json:"code,omitempty" binding:"omitempty"`
	Message string      `json:"message,omitempty" binding:"omitempty"`
	Data    interface{} `json:"data"`
}

func ResError(code int, data interface{}) ResponseError {
	return ResponseError{
		Code:    code,
		Message: "false",
		Data: struct {
			Message string `json:"message,omitempty" binding:"omitempty"`
		}{
			Message: data.(string),
		},
	}
}

func ResSuccess(code int, data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Code:    code,
		Message: "success",
		Data:    data,
	}
}
