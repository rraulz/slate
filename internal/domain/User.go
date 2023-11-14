package domain

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

func NewUser(id int, username, password, email string) *User {

	return &User{id, username, password, email}
}
