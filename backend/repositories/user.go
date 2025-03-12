package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/nambuitechx/go-social/models"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{ DB: db }
}

func (r *UserRepository) SelectUsers(limit *int, offset *int) ([]models.UserModel, error) {
	users := []models.UserModel{}
	var err error

	if *limit < 0 {
		statement := "SELECT id, email, created_at, updated_at FROM users"
		err = r.DB.Select(&users, statement)
	} else {
		statement := "SELECT id, email, created_at, updated_at FROM users LIMIT $1 OFFSET $2"
		err = r.DB.Select(&users, statement, *limit, *offset)
	}
	
	return users, err
}

func (r *UserRepository) SelectUserById(id *string) (*models.UserModel, error) {
	user := &models.UserModel{}

	statement := "SELECT id, email, created_at, updated_at FROM users WHERE id = $1"
	err := r.DB.Get(user, statement, *id)
	return user, err
}

func (r *UserRepository) SelectUserByEmail(email *string, selectPassword bool) (*models.UserModel, error) {
	user := &models.UserModel{}
	var statement string
	if !selectPassword {
		statement = "SELECT id, email, created_at, updated_at FROM users WHERE email = $1"
	} else if selectPassword {
		statement = "SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1"
	}
	err := r.DB.Get(user, statement, *email)
	return user, err
}

func (r *UserRepository) InsertUser(payload *models.CreateUserPayload) (*models.UserModel, error) {
	var user = models.UserModel{}
	statement := "INSERT INTO users(id, email, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id, email, created_at, updated_at"
	id := uuid.New().String()
	now := time.Now().Unix()
	err := r.DB.Get(
		&user,
		statement,
		id,
		payload.Email,
		payload.Password,
		now,
		now,
	)
	return &user, err
}

func (r *UserRepository) DeleteUserById(id *string) error {
	statement := "DELETE FROM users WHERE id = $1"
	_, err := r.DB.Exec(statement, *id)
	return err
}

func (r *UserRepository) DeleteUserByEmail(email *string) error {
	statement := "DELETE FROM users WHERE email = $1"
	_, err := r.DB.Exec(statement, *email)
	return err
}
