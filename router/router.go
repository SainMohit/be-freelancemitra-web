package router

import (
	"github.com/be-freelancemitra-web/src/api"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(obj api.ApiV1_0Layer) (router *gin.Engine) {
	router = gin.New()
	router.Use(gin.Recovery())
	router.GET("/health", obj.GetV1_0ApiLayer().Health)
	return router
}
