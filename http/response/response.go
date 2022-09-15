package response

import (
	"errors"
	"net/http"

	"github.com/herryg91/kless_go_lib/http/middleware"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success     bool        `json:"success"`
	ProcessTime string      `json:"process_time"`
	ErrorCode   string      `json:"error_code,omitempty"`
	Message     string      `json:"message,omitempty"`
	Data        interface{} `json:"data"`
}

func (r *Response) SetProcessTime(c *gin.Context) {
	r.ProcessTime = middleware.GetProcessTime(c)
}

func Success(c *gin.Context, status int, data interface{}) {
	if status >= 400 || status < 200 {
		c.AbortWithError(http.StatusInternalServerError, errors.New("success status must be on the range 200-399"))
		return
	}

	resp := Response{
		Success:     true,
		ProcessTime: middleware.GetProcessTime(c),
		Data:        data,
	}
	c.JSON(http.StatusNotFound, resp)
}

func Error(c *gin.Context, status int, err_code string, err_message string, detail_errors ...ErrorFieldResponse) {
	if status < 400 {
		c.AbortWithError(http.StatusInternalServerError, errors.New("http status must be < 400"))
		return
	}

	resp := Response{
		Success:     false,
		Message:     err_message,
		ErrorCode:   err_code,
		ProcessTime: middleware.GetProcessTime(c),
		Data:        detail_errors,
	}
	c.JSON(http.StatusNotFound, resp)
}
