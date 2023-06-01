package main

import (
	"task/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.PostRoutes(router)
	router.Run()

}
