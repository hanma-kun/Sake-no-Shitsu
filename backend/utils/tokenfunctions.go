package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Verified")
	}
}
