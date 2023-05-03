package middleware

import (
	"fmt"
	"net/http"
	"sake/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		if strings.HasPrefix(c.Request.URL.Path, "/cart") {

			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "No Authorization Header Provided"})
				c.Abort()
				return
			}
			claims, err := utils.ValidateToken(token)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			c.Set("uid", claims.Uid)
			c.Next()
		} else {
			c.Next()
		}
	}
}
