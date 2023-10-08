package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrorInvalidID   = errors.New("invalid id provided")
	ErrorNameTooLong = errors.New("invalid name provided")
)

type User struct {
	ID       ID
	Username Username
	Email    Email
}

type ID string
type Email string
type Username string

func New(email Email, username Username) User {
	return User{
		ID:       ID(uuid.NewString()),
		Email:    email,
		Username: username,
	}
}

func NewID(i string) (ID, error) {
	//domain logic
	// _, err := uuid.Parse(i)
	// if err != nil {
	// 	return errors.Join(err, ErrorInvalidID)
	// }

	return ID(i), nil
}

func (i ID) String() string {
	return string(i)
}

func NewEmail(e string) (Email, error) {
	//domain logic
	return Email(e), nil
}

func (e Email) String() string {
	return string(e)
}

func NewUsername(u string) (Username, error) {
	//domain logic
	return Username(u), nil
}

func (u Username) String() string {
	return string(u)
}
