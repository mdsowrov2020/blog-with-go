package domain

type User struct {
	ID       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	IsAuthor bool   `json:"is_author" db:"is_author"`
}
