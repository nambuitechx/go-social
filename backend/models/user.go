package models

import "time"

type UserModel struct {
	ID			string		`db:"id"`
	Email		string		`db:"email"`
	CreatedAt	time.Time	`db:"created_at"`
	UpdatedAt	time.Time	`db:"updated_at"`
}

type GetUserQuery struct {
	Limit int	`form:"limit" default:"10"`
	Offset int	`form:"offset" default:"0"`
}

type GetUserParam struct {
	ID string	`uri:"id"`
}

type CreateUserPayload struct {
	Email	string	`json:"email" binding:"required"`
}
