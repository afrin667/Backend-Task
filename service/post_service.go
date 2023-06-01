package service

import (
	"task/entity"
)

type PostService interface {
	Save(entity.Post) entity.Post
	FindAll() []entity.Post
	LikePost(entity.Request) entity.Post
	UnlikePost(entity.Request) entity.Post
	AddComment(entity.Request) entity.Post
}

type postService struct {
	posts  []entity.Post
	postId int
}

func NewPostService() PostService {
	return &postService{
		postId: 1,
	}
}

func (s *postService) Save(post entity.Post) entity.Post {
	newPost := &entity.Post{
		Post_Id:  s.postId,
		Content:  post.Content,
		MediaUrl: post.MediaUrl,
	}
	s.posts = append(s.posts, *newPost)
	s.postId++
	return *newPost
}

func (s *postService) FindAll() []entity.Post {

	return s.posts
}

func (s *postService) LikePost(request entity.Request) entity.Post {
	var post entity.Post
	for i := 0; i < len(s.posts); i++ {
		if s.posts[i].Post_Id == request.Id {
			s.posts[i].IsLiked = true
			post = s.posts[i]
		}
	}

	return post
}

func (s *postService) UnlikePost(request entity.Request) entity.Post {
	var post entity.Post
	for i := 0; i < len(s.posts); i++ {
		if s.posts[i].Post_Id == request.Id {
			s.posts[i].IsLiked = false
			post = s.posts[i]
		}
	}

	return post
}

func (s *postService) AddComment(request entity.Request) entity.Post {
	var post entity.Post
	var comments []string

	for i := 0; i < len(s.posts); i++ {
		if s.posts[i].Post_Id == request.Id {
			comments = append(s.posts[i].Comment, "hello")
			s.posts[i].Comment = comments
			post = s.posts[i]
		}
	}
	return post
}
