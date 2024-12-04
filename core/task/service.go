package task

import (
	"time"

	"github.com/juheth/to-do/core/models"
)

type (
	Service interface {
		Create(name, descrip, userID string, dueDate time.Time) error
		GetUserById(id string) error
		GetAllTask() ([]models.Task, error)
		GetAllTaskById(id string) ([]models.Task, error)
		GetTaskById(id string) (models.Task, error)
		UpDateTask(id, name, descrip, userID string, dueDate, create time.Time, status bool) (string, error)
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

func (s service) Create(name, descrip, userID string, Date time.Time) error {

	task := &models.Task{
		Id:        "",
		Name:      name,
		Date:      Date,
		UserID:    userID,
		State:     false,
		Create_at: time.Now(),
		Update_at: time.Now(),
	}
	err := s.repo.Create(task)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetUserById(id string) error {
	err := s.repo.GetUserById(id)
	if err != nil {
		return err
	}
	return nil
}
func (s service) GetAllTask() ([]models.Task, error) {
	tasks, err := s.repo.GetAllTask()
	if err != nil {
		return tasks, nil
	}
	return tasks, nil
}

func (s service) GetTaskById(id string) (models.Task, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return task, err
	}
	return task, nil
}
func (s service) GetAllTaskById(id string) ([]models.Task, error) {
	tasks, err := s.repo.GetAllTaskById(id)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s service) UpDateTask(id, name, descrip, userID string, Date, create time.Time, status bool) (string, error) {
	task := &models.Task{
		Id:        id,
		Name:      name,
		Date:      Date,
		UserID:    userID,
		State:     status,
		Create_at: create,
		Update_at: time.Now(),
	}

	err := s.repo.UpDateTask(task)
	if err != nil {
		return "", err
	}
	update := task.Update_at.Format("2006-01-02 15:04:05")
	return update, err
}
