package entity

type Request struct {
	Id         int    `json:"id"`
	NewComment string `json:"newComment"`
}
