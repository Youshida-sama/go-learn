package repositories

import (
	"main/models"
	"main/storage"
	"time"
)

// Создает запись о действии пользователя
func CreateUserAction(userAction models.UserAction) (models.UserAction, error) {
	db := storage.GetDB()
	sqlStatement := `call create_user_action($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, userAction.ID, userAction.Name, userAction.Surname, userAction.Action, userAction.Time.Format(time.RFC3339))

	if err != nil {
		return userAction, err
	}

	return userAction, nil
}
