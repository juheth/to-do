package proyect

import "github.com/gin-gonic/gin"

type (
	Controller func(c *gin.Context)
	EndPoints  struct {
	}
)
