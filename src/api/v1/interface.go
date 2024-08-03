package v1_0

import "github.com/gin-gonic/gin"

type ApiLayer interface {
	Health(c *gin.Context)
}

type apiObject struct {
}

func NewApiObj() ApiLayer {

	return &apiObject{}
}
