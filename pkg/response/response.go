package response

import (
	"net/http"

	errors "authorizer/pkg/error"
	"github.com/gin-gonic/gin"
)

type Error struct {
	GlobalError string            `json:"GlobalError,omitempty"`
	FieldsError map[string]string `json:"FieldsError,omitempty"`
	Result      *interface{}      `json:"Result,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, data)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func BadRequest(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{})
}

func NotAuthorized(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, err.Error())
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{})
}

func ErrorResponseGlobal(c *gin.Context, globalError interface{}, data interface{}) {
	result := &Error{}

	if data != nil {
		result.Result = &data
	}

	err, ok := globalError.(*errors.PermissionError)
	if ok {
		result.GlobalError = err.Error()
		c.AbortWithStatusJSON(http.StatusForbidden, result)

		return
	}

	err1, ok1 := globalError.(error)
	if ok1 {
		result.GlobalError = err1.Error()
		c.JSON(http.StatusBadRequest, result)

		return
	}

	result.GlobalError = globalError.(string)

	c.JSON(http.StatusBadRequest, result)
}

func ErrorResponseFields(c *gin.Context, fieldsError errors.FieldErrors, data interface{}) {
	result := &Error{
		FieldsError: fieldsError,
	}

	if data != nil {
		result.Result = &data
	}

	c.JSON(http.StatusBadRequest, result)
}
