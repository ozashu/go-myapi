package models

import (
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id`
	ArticleID int       `json:"article_id`
	Message   string    `json:message`
	CreatedAt time.Time `json:created_at`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"titile"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_name"`
	CommentList []Comment `json:"comment_list"`
	CreatedAt   time.Time `json:"created_at"`
}
