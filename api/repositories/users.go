package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        string         `db:"id"`
	Email     string         `db:"email"`
	Credits   int            `db:"credits"`
	PublicKey sql.NullString `db:"public_key"`
	OAuthID   string         `db:"oauth_id"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
}

type UsersRepository struct {
	DB *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{DB: db}
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
		INSERT INTO users (email, oauth_id, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
	`
	_, err := r.DB.Exec(query, user.Email, user.OAuthID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepository) UpdatePublicKey(id string, publicKey string) error {
	query := `
		UPDATE users SET public_key = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2;
	`
	_, err := r.DB.Exec(query, publicKey, id)
	if err != nil {
		return err
	}
	return nil
}
