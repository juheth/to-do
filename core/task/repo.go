package task

import (
	"github.com/juheth/to-do/core/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(task *models.Task) error
		GetUserById(id string) error
		GetTaskById(id string) (models.Task, error)
		GetAllTaskById(id string) ([]models.Task, error)
		GetAllTask() ([]models.Task, error)
		UpDateTask(task *models.Task) error
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

func (repo *repo) Create(task *models.Task) error {
	if err := repo.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUserById(id string) error {
	user := models.User{}
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (repo *repo) GetTaskById(id string) (models.Task, error) {
	task := models.Task{}
	err := repo.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}
func (repo *repo) GetAllTaskById(id string) ([]models.Task, error) {
	var tasks []models.Task
	err := repo.db.Where("user_id = ?", id).Find(&tasks)
	if err.Error != nil {
		return nil, err.Error
	}
	return tasks, nil
}
func (repo *repo) GetAllTask() ([]models.Task, error) {
	var tasks []models.Task
	err := repo.db.Find(&tasks)
	if err.Error != nil {
		return nil, err.Error
	}
	return tasks, nil
}
func (repo *repo) UpDateTask(task *models.Task) error {
	if err := repo.db.Save(task).Error; err != nil {
		return err
	}
	return nil
}