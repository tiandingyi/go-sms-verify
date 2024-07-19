package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// validate.Struct 通常情况下，它用于检查给定的数据结构体是否符合特定的验证规则。
// type User struct {
//    Username string `validate:"required,min=3,max=20"`

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	if err := validate.Struct(&data); err != nil {
		return err
	}
	return nil
}

func (app *Config) writeJSON(c *gin.Context, status int, data any) {
	c.JSON(status, jsonResponse{Status: status, Message: "success", Data: data})
}

func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})

}
