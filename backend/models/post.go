package models

import "time"

type PostModel struct {
	ID			string		`db:"id" json:"id"`
	Content		string		`db:"content" json:"content"`
	UserId		string		`db:"user_id" json:"userId"`
	CreatedAt	time.Time	`db:"created_at" json:"createdAt"`
	UpdatedAt	time.Time	`db:"updated_at" json:"updatedAt"`
}

type GetPostQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetPostParam struct {
	ID string	`uri:"id"`
}

type CreatePostPayload struct {
	Content	string	`json:"content" binding:"required"`
}
