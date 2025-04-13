package repositories

import (
	"database/sql"
	"time"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (email, password, username, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`

	now := time.Now()
	err = r.db.QueryRow(
		query,
		user.Email,
		string(hashedPassword),
		user.Username,
		now,
		now,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	user.CreatedAt = now
	user.UpdatedAt = now
	user.Password = ""

	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password, username, created_at, updated_at 
              FROM users WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) VerifyPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
