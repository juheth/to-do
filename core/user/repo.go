package user

import (
	"github.com/juheth/to-do/core/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Register(user *models.User) error
		GetUser() ([]models.User, error)
		GetUserByMail(email string) (models.User, error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (repo *repo) Register(user *models.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser() ([]models.User, error) {
	var user []models.User

	tx := repo.db.Select("id", "name", "email", "password").Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (repo *repo) GetUserByMail(email string) (models.User, error) {
	user := models.User{}
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
