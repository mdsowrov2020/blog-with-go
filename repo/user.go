package repo

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAuthor bool   `json:"is_author"`
}

type UserRepo interface {
	Create(p User) (*User, error)
	Find(email, password string) (*User, error)
	// List() ([]*User, error)
	// Update(p User) (*User, error)
	// Delete(id int) error
}

type userRepo struct {
	userList []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.userList) + 1
	r.userList = append(r.userList, user)
	return &user, nil
}

func (r userRepo) Find(email string, password string) (*User, error) {
	for _, user := range r.userList {
		if user.Email == email && user.Password == password {
			return &user, nil
		}
	}

	return nil, nil
}
