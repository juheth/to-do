package models

type Proyect struct {
	id     int
	name   string `json:"name"`
	userID string `json:"id_user"`
}
