package models

import "time"

type Task struct {
	Id        string
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	State     bool      `json:"state"`
	UserID    string    `json:"id_user"`
	ProyectID string    `json:"id_proyect"`
	Create_at time.Time `json:"create_at"`
	Update_at time.Time `json:"update_at"`
}
