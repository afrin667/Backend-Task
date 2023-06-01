package entity

type Post struct {
	Post_Id   int      `json:"postId"`
	Content   string   `json:"content"`
	MediaUrl  string   `json:"mediaUrl"`
	IsLiked   bool     `json:"isLiked"`
	LikeCount int      `json:"likeCount"`
	Comment   []string `json:"comments"`
}
