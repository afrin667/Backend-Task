package controller

import (
	"task/entity"
	"task/service"

	"github.com/gin-gonic/gin"
)

type PostController interface {
	FindAll() []entity.Post
	Save(ctx *gin.Context) entity.Post
	LikePost(ctx *gin.Context) entity.Post
	UnlikePost(ctx *gin.Context) entity.Post
	AddComment(ctx *gin.Context) entity.Post
}

type postController struct {
	service service.PostService
}

func NewPostController(service service.PostService) PostController {
	return &postController{
		service: service,
	}
}

func (c *postController) FindAll() []entity.Post {
	return c.service.FindAll()
}

func (c *postController) Save(ctx *gin.Context) entity.Post {
	var post entity.Post
	ctx.BindJSON(&post)
	var createdPost = c.service.Save(post)
	return createdPost
}

func (c *postController) LikePost(ctx *gin.Context) entity.Post {
	var request entity.Request
	ctx.BindJSON(&request)
	var post = c.service.LikePost(request)
	return post
}

func (c *postController) UnlikePost(ctx *gin.Context) entity.Post {
	var request entity.Request
	ctx.BindJSON(&request)
	var post = c.service.UnlikePost(request)
	return post
}

func (c *postController) AddComment(ctx *gin.Context) entity.Post {
	var request entity.Request
	ctx.BindJSON(&request)
	var post = c.service.AddComment(request)
	return post
}
