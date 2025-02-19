package models

import "time"

type PostModel struct {
	ID			string		`db:"id"`
	Content		string		`db:"content"`
	UserId		string		`db:"user_id"`
	CreatedAt	time.Time	`db:"created_at"`
	UpdatedAt	time.Time	`db:"updated_at"`
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
