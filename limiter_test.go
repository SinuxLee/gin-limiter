package ginlimiter

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRateLimiter(t *testing.T) {
	r := gin.Default()
	r.Use(NewRateLimiter(time.Second, 5000, func(ctx *gin.Context) (string, error) {
		return "", nil
	}).Middleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	r.Run(":8086")
}
