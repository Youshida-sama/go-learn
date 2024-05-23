package requests

import (
	"main/enums"
	"main/models"
	"time"
)

// Запрос на логирование действий пользователя
type UserActionRequest struct {
	ID      int          `json:"id" validate:"required,gte=0"`
	Name    string       `json:"name" validate:"req=Surname"`
	Surname string       `json:"surname" validate:"req=Name"`
	Action  enums.Action `json:"action" validate:"required,enum"`
	Time    string       `json:"time" validate:"isoTime"`
}

// Конвертация запроса в бизнес-модель
func (uar UserActionRequest) Map() (ua *models.UserAction, err error) {
	t, err := time.Parse(time.RFC3339, uar.Time)

	if err != nil {
		return
	}

	ua = &models.UserAction{
		ID:      uar.ID,
		Name:    uar.Name,
		Surname: uar.Surname,
		Time:    t,
	}

	return
}
