package models

import "time"

type UserModel struct {
	ID			string		`db:"id" json:"id"`
	Email		string		`db:"email" json:"email"`
	CreatedAt	time.Time	`db:"created_at" json:"createdAt"`
	UpdatedAt	time.Time	`db:"updated_at" json:"updatedAt"`
}

type GetUserQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetUserByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type CreateUserPayload struct {
	Email	string	`json:"email" binding:"required"`
}
