package user

import (
	"regexp"

	"github.com/juheth/to-do/core/models"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service interface {
		RegisterUser(name, email, password string) error
		IsValidMail(email string) bool
		GetUser() ([]models.User, error)
		GetUserByMail(email string) (models.User, error)
		ValidPassword(email, password string) (bool, error)
	}
	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) RegisterUser(name, email, password string) error {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return s.repo.Register(&user)
}

func (s service) IsValidMail(email string) bool {
	validMail := regexp.MustCompile("^[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})$")
	return validMail.MatchString(email)
}

func (s service) GetUser() ([]models.User, error) {
	users, err := s.repo.GetUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) GetUserByMail(email string) (models.User, error) {
	user, err := s.repo.GetUserByMail(email)
	return user, err
}

func (s service) ValidPassword(email, password string) (bool, error) {
	user, err := s.repo.GetUserByMail(email)
	if err != nil {
		return true, err
	}

	passwordByte := []byte(password)
	passwordDB := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(passwordDB, passwordByte)

	if err != nil {
		return false, err
	}

	return true, nil
}
