package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	IsAuthor bool   `json:"is_author" db:"is_author"`
}

type UserRepo interface {
	Create(p User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Update(p User) (*User, error)
	// Delete(id int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
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

	// 3. Scan the returned ID into the struct
	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID = userID

	return &user, nil
}

func (r *userRepo) Find(email string, password string) (*User, error) {
	// for _, user := range r.userList {
	// 	if user.Email == email && user.Password == password {
	// 		return &user, nil
	// 	}
	// }

	// return nil, nil

	var user User

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
