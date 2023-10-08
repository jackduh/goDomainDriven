package service

import (
	"context"

	"github.com/jackduh/goDomainDriven/domain/user"
)

type Service struct {
}

// DTO (Data Transfer Object) hold primative fields
type Request struct {
	Username string
	Email    string
}

type SingUpResponse struct {
	ID string
}

func (d *Service) SignUp(ctx context.Context, username string, email string) (*SingUpResponse, error) {
	//Get request fields to speak to our domain language
	userName, err := user.NewUsername(username)

	if err != nil {
		return nil, err
	}

	userEmail, err := user.NewEmail(email)

	if err != nil {
		return nil, err
	}

	u := user.New(userEmail, userName)

	// pass to repos

	return &SingUpResponse{
		ID: u.ID.String(),
	}, nil
}
