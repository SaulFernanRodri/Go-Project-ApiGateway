package main

import (
	"api-gateway/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.AuthRoutes(router)
	routers.UserRoutes(router)

	router.Run(":8080")
}
