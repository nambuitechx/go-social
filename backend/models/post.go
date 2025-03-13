package models

type PostModel struct {
	ID			string		`db:"id" json:"id"`
	Content		string		`db:"content" json:"content"`
	UserId		string		`db:"user_id" json:"userId"`
	CreatedAt	int			`db:"created_at" json:"createdAt"`
	UpdatedAt	int			`db:"updated_at" json:"updatedAt"`
}

type GetPostQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetPostByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type CreatePostPayload struct {
	Content	string	`json:"content" binding:"required"`
}
