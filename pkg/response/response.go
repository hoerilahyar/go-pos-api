package response

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success response
func Success(c *gin.Context, message string, data ...interface{}) {
	resp := Response{
		Status:  "success",
		Message: message,
	}

	if len(data) > 0 {
		resp.Data = data[0]
	}

	c.JSON(http.StatusOK, resp)
}

// Error response
func Error(c *gin.Context, err error, data ...interface{}) {
	var errMessage string

	switch {
	case errors.Is(err, io.EOF):
		errMessage = "Request body is empty or invalid JSON format"
	case err != nil:
		errMessage = err.Error()
	default:
		errMessage = "Unexpected error"
	}

	resp := Response{
		Status:  "error",
		Message: errMessage,
	}

	if len(data) > 0 {
		resp.Data = data[0]
	}

	c.JSON(http.StatusBadRequest, resp)
}
