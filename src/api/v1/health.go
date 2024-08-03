package v1_0

import (
	"github.com/gin-gonic/gin"
)

func (s *apiObject) Health(c *gin.Context) {
	c.String(200, "I am healthy")
}
