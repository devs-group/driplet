package repositories

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Credits   int    `db:"credits"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type UsersRepository struct {
	DB *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

func (r *UsersRepository) FindByID(id int64) (*User, error) {
	var user User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepository) Create(user *User) error {
	query := `
		INSERT INTO users (email, created_at, updated_at)
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, created_at, updated_at
	`
	return r.DB.QueryRowx(query, user.Email).StructScan(user)
}
