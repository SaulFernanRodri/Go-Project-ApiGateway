package routers

import (
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/auth/register", func(c *gin.Context) {
		proxy.ProxyRequest(c, "http://localhost:8081/auth/register")
	})
	router.POST("/auth/login", func(c *gin.Context) {
		proxy.ProxyRequest(c, "http://localhost:8081/auth/login")
	})
}
