package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juheth/to-do/core/jwt"
)

func ValidToken(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token Requerido"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	isValidToken, err := jwt.Token(token)
	if !isValidToken {
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "Token invalid",
				"err":     err.Error(),
			})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	c.Next()
}
