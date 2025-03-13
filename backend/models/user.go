package models

type UserModel struct {
	ID			string		`db:"id" json:"id"`
	Email		string		`db:"email" json:"email"`
	Password	string		`db:"password" json:"password"`
	CreatedAt	int			`db:"created_at" json:"createdAt"`
	UpdatedAt	int			`db:"updated_at" json:"updatedAt"`
}

type GetUserQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetUserByIdParam struct {
	ID		string		`uri:"id" binding:"required"`
}

type GetUserByEmailParam struct {
	Email	string		`uri:"email" binding:"required"`
}

type CreateUserPayload struct {
	Email		string	`json:"email" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

type TokenInfo struct {
	Token		string	`json:"token"`
}

type UserInfo struct {
	ID			string	`json:"id"`
	Email		string	`json:"email"`
}
