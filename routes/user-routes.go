package routes

import (
	"api-gateway/middlewares"
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	view := router.Group("/view")
	{
		view.GET("/", func(c *gin.Context) {
			proxy.ProxyRequest(c, "http://localhost:8082/view")
		})
	}

	secured := router.Group("/users")
	secured.Use(middlewares.AuthJWT())
	{
		secured.GET("/", func(c *gin.Context) {
			proxy.ProxyRequest(c, "http://localhost:8082/users")
		})
		secured.POST("/", func(c *gin.Context) {
			proxy.ProxyRequest(c, "http://localhost:8082/users")
		})
		secured.PUT("/:id", func(c *gin.Context) {
			proxy.ProxyRequest(c, "http://localhost:8082/users/"+c.Param("id"))
		})
		secured.DELETE("/:id", func(c *gin.Context) {
			proxy.ProxyRequest(c, "http://localhost:8082/users/"+c.Param("id"))
		})
	}
}
