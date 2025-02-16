package types

import "time"

type UserStore interface { // used cause we can easily test interfaces in go
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

// type mockUserStore struct{}

// func GetUserByEmail(email string) (*User, error) { // easily create tests using these methods
// 	return nil, nil
// }

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=4,max=30"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=30"`
}
