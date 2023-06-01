package main

import (
	"fmt"
	"task/controller"
	"task/service"

	"github.com/gin-gonic/gin"
)

var (
	// userService    service.UserService       = service.NewUserService()
	// userController controller.UserController = controller.NewUserController(userService)

	postService    service.PostService       = service.NewPostService()
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	fmt.Println("Hello")

	router := gin.Default()

	// router.GET("/users", func(ctx *gin.Context) {
	// 	ctx.JSON(200, userController.FindAll())
	// })

	// router.POST("/users", func(ctx *gin.Context) {
	// 	ctx.JSON(200, userController.Save(ctx))
	// })

	router.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, postController.FindAll())
	})

	router.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, postController.Save(ctx))
	})

	router.POST("/like", func(ctx *gin.Context) {
		ctx.JSON(200, postController.LikePost(ctx))
	})

	router.POST("/unlike", func(ctx *gin.Context) {
		ctx.JSON(200, postController.UnlikePost(ctx))
	})

	router.Run(":8080")
}
