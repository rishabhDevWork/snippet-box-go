package mocks

import (
	"github.com/rishabhdevwork/snippet-box-go/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name string, email string, password string) error {

	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email string, password string) (int, error) {
	if email == "rishabh@example.com" && password == "password" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil

	}
}
