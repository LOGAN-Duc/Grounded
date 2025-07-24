package common

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginationData struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

type ResponseDataWithTotal struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
}

func ResponseGinWithValidationError(c *gin.Context, err error) {
	if newErr, ok := err.(validator.ValidationErrors); ok {
		errMap := make(map[string]string)
		for _, err := range newErr {
			field := strings.Split(err.StructField(), ".")[0]
			if len(field) > 0 {
				field = strings.ToLower(field[:1]) + field[1:]
				field = strings.ReplaceAll(field, "ID", "Id")
				if err.Tag() == "required" {
					errMap[field] = field + " is required"
				} else {
					errMap[field] = field + " must be a valid " + err.Tag()
				}
			}
		}
		c.JSON(http.StatusBadRequest, ResponseData{
			Status:  http.StatusBadRequest,
			Message: errMap,
		})
	}
	if strings.Contains(err.Error(), "parsing time") {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status:  http.StatusBadRequest,
			Message: "Time invalid",
		})
	}
	if strings.Contains(err.Error(), "invalid") {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
}

func ResponseGin(c *gin.Context, data interface{}, message ...string) {
	c.JSON(http.StatusOK, ResponseData{
		Status:  http.StatusOK,
		Message: strings.Join(message, ", "),
		Data:    data,
	})
}

func ResponseGinV1(c *gin.Context, data interface{}, total int, message ...string) {
	c.JSON(http.StatusOK, ResponseDataWithTotal{
		Status:  http.StatusOK,
		Message: strings.Join(message, ", "),
		Data:    data,
		Total:   total,
	})
}

func ResponseGinWithError(c *gin.Context, message ...string) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Status:  http.StatusBadRequest,
		Message: strings.Join(message, ", "),
	})
}

func ResponseGinWithErrorAndInfo(c *gin.Context, data interface{}, message ...string) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Status:  http.StatusBadRequest,
		Message: strings.Join(message, ", "),
		Data:    data,
	})
}

func ResponseGinWithNotFound(c *gin.Context, message ...string) {
	c.JSON(http.StatusNotFound, ResponseData{
		Status:  http.StatusNotFound,
		Message: strings.Join(message, ", "),
	})
}

func ResponseGinWithForbidden(c *gin.Context, message ...string) {
	c.JSON(http.StatusForbidden, ResponseData{
		Status:  http.StatusForbidden,
		Message: strings.Join(message, ", "),
	})
}

func ResponseGinWithUnauthorized(c *gin.Context, message ...string) {
	c.JSON(http.StatusUnauthorized, ResponseData{
		Status:  http.StatusUnauthorized,
		Message: strings.Join(message, ", "),
	})
}

func ResponseGinWithInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ResponseData{
		Status:  http.StatusInternalServerError,
		Message: "An error appeared! Please try again",
	})
}

func ResponseGinWithPagination(c *gin.Context, data interface{}, total int64, message ...string) {
	c.JSON(http.StatusOK, ResponseData{
		Status:  http.StatusOK,
		Message: strings.Join(message, ", "),
		Data: PaginationData{
			Data:  data,
			Total: total,
		},
	})
}

func ResponseWithErrorAndData(c *gin.Context, data interface{}, message ...string) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Status:  http.StatusBadRequest,
		Data:    data,
		Message: strings.Join(message, ", "),
	})
}

func ResponseWithError(c *gin.Context, errCode int, message ...string) {
	c.JSON(errCode, ResponseData{
		Status:  errCode,
		Message: strings.Join(message, ", "),
	})
}

func ResponseGinWithCursor(c *gin.Context, data interface{}, paging interface{}, filter interface{}, message ...string) {
	c.JSON(http.StatusOK, gin.H{
		"result":     data,
		"pagingData": paging,
		"filter":     filter,
		"message":    strings.Join(message, ", "),
	})
}
