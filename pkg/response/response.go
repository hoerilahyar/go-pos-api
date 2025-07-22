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
type PaginatedResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Meta    PaginateMeta
}

type PaginateMeta struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalPages int   `json:"total_pages"`
	TotalRows  int64 `json:"total_rows"`
}

// Paginate Response
func Paginated(c *gin.Context, message string, data interface{}, page, limit int, totalRows int64) {
	totalPages := int((totalRows + int64(limit) - 1) / int64(limit))

	meta := PaginateMeta{

		Limit:      limit,
		Page:       page,
		TotalPages: totalPages,
		TotalRows:  totalRows,
	}

	resp := PaginatedResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	c.JSON(http.StatusOK, resp)
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
