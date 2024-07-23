package models

import "time"

type Task struct {
	Id        string
	Name      string    `json:"name"`
	date      time.Time `json:"date"`
	State     bool      `json:"state"`
	ProyectID string    `json:"id_proyect"`
	Create_at time.Time `json:"create_at"`
	Update_at time.Time `json:"update_at"`
}
