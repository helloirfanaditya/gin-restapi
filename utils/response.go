package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Code    int    `json:"code,omitempty" binding:"omitempty"`
	Message string `json:"message,omitempty" binding:"omitempty"`
	Data    struct {
		Message string `json:"message,omitempty" binding:"omitempty"`
	} `json:"data,omitempty" binding:"omitempty"`
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

func ResSuccess(code int, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"data": data,
	}
}
