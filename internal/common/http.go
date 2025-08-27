package common

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrType int

const (
	InternalServerErr ErrType = 1
	ValidationErr     ErrType = 2
	PostgresErr       ErrType = 3
	OtherErr          ErrType = 4
)

func ErrValidation(c *gin.Context, err error) {
	if newErr, ok := err.(validator.ValidationErrors); ok {
		errMap := make(map[string]interface{})
		for _, err := range newErr {
			fields := strings.Split(err.Namespace(), ".")
			lastField := formatField(fields[len(fields)-1])
			message := ""
			tagsMap := map[string]string{
				"eq":  "equal",
				"gt":  "greater than",
				"gte": "greater than or equal",
				"lt":  "less than",
				"lte": "less than or equal",
				"ne":  "not equat",
			}
			switch err.Tag() {
			case "required":
				message = lastField + " is required"
			case "eq", "gt", "gte", "lt", "lte", "ne":
				message = fmt.Sprintf("%s must be %s %s", lastField, tagsMap[err.Tag()], err.Param())
			default:
				message = lastField + " must be a valid " + err.Tag()
			}
			if len(fields) == 1 {
				errMap[lastField] = message
			} else {
				data := map[string]interface{}{
					lastField: message,
				}
				for i := len(fields) - 2; i > 0; i-- {
					if strings.Contains(fields[i], "[") {
						indStringArr := strings.Split(fields[i], "[")
						indString := strings.ReplaceAll(indStringArr[1], "]", "")
						ind, _ := strconv.Atoi(indString)
						data = map[string]interface{}{
							formatField(indStringArr[0]): map[int]interface{}{
								ind: data,
							},
						}
					} else {
						data = map[string]interface{}{
							formatField(fields[i]): data,
						}
					}
				}
				errMap = data
			}
		}
		c.JSON(http.StatusBadRequest, ResponseData{
			Status:  http.StatusBadRequest,
			Message: "data invalid",
			Errors:  errMap,
		})
		return
	}
	if strings.Contains(err.Error(), "invalid UUID") {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status:  http.StatusBadRequest,
			Message: "UUID invalid",
		})
		return
	}
	if strings.Contains(err.Error(), "parsing time") {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status:  http.StatusBadRequest,
			Message: "Time invalid",
		})
		return
	}
	if strings.Contains(err.Error(), "cannot unmarshal string into") {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status:  http.StatusBadRequest,
			Message: "invalid json field format",
		})
		return
	}
	c.JSON(http.StatusBadRequest, ResponseData{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	})
}

func formatField(field string) string {
	formattedField := strings.ReplaceAll(field, "ID", "Id")
	formattedField = strings.ToLower(formattedField[:1]) + formattedField[1:]
	return formattedField
}

func Success(c *gin.Context, data interface{}, message ...string) {
	c.JSON(http.StatusOK, ResponseData{
		Status:  http.StatusOK,
		Message: strings.Join(message, ", "),
		Data:    data,
	})
}

func SuccessPaginate(c *gin.Context, data interface{}, total int64, message ...string) {
	c.JSON(http.StatusOK, ResponseData{
		Status:  http.StatusOK,
		Message: strings.Join(message, ", "),
		Data: PaginationData{
			Data:  data,
			Total: total,
		},
	})
}
