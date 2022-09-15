package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

const PROCESS_TIME_CONTEXT = "x-process-time"

func UseProcessTimeMiddleware(c *gin.Context) {
	c.Set(PROCESS_TIME_CONTEXT, time.Now())
}

func GetProcessTime(c *gin.Context) string {
	return time.Since(c.GetTime(PROCESS_TIME_CONTEXT)).String()
}
