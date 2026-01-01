package repo

import (
	"database/sql"
	"fmt"

	"blog/domain"
	"blog/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
		full_name,
		email,
		password,
		is_author
		)
		VALUES (
		:full_name,
		:email,
		:password,
		:is_author
		)
		RETURNING id
		`

	var userID int

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID = userID

	return &user, nil
}

func (r *userRepo) Find(email string, password string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, full_name, email, password, is_author
		FROM
		users
		WHERE email = $1 AND password = $2
	`

	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
