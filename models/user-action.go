package models

import (
	"main/enums"
	"time"
)

// Бизнес-модель действия пользователя
type UserAction struct {
	ID      int
	Name    string
	Surname string
	Action  enums.Action
	Time    time.Time
}
