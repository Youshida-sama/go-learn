package models

import "time"

//Бизнес-модель действия пользователя
type UserAction struct {
	ID      int
	Name    string
	Surname string
	Time    time.Time
}
