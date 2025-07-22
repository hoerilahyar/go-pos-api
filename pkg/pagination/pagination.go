package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

// Default limit if not set
const defaultLimit = 10

func Paginate(c *gin.Context, db *gorm.DB, model interface{}, out interface{}) (*Pagination, error) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", strconv.Itoa(defaultLimit))

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset := (page - 1) * limit

	var totalRows int64
	if err := db.Model(model).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if err := db.Limit(limit).Offset(offset).Find(out).Error; err != nil {
		return nil, err
	}

	totalPages := int((totalRows + int64(limit) - 1) / int64(limit)) // pembulatan ke atas

	return &Pagination{
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       out,
	}, nil
}
