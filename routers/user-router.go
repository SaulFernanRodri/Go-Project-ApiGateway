package routers

import (
	"api-gateway/middlewares"
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")
	{
		view := api.Group("/view")
		{
			view.GET("/", func(c *gin.Context) {
				proxy.ProxyRequest(c, "http://localhost:8082/api/v1/view")
			})
		}

		secured := api.Group("/users")
		secured.Use(middlewares.AuthJWT())
		{
			secured.GET("/", func(c *gin.Context) {
				proxy.ProxyRequest(c, "http://localhost:8082/api/v1/users")
			})
			secured.POST("/", func(c *gin.Context) {
				proxy.ProxyRequest(c, "http://localhost:8082/api/v1/users")
			})
			secured.PUT("/:id", func(c *gin.Context) {
				proxy.ProxyRequest(c, "http://localhost:8082/api/v1/users/"+c.Param("id"))
			})
			secured.DELETE("/:id", func(c *gin.Context) {
				proxy.ProxyRequest(c, "http://localhost:8082/api/v1/users/"+c.Param("id"))
			})
		}
	}
}
