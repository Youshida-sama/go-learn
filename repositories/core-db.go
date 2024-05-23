package repositories

import (
	"context"
	"main/logging"
	"main/models"
	"main/storage"
	"time"
)

// Создает запись о действии пользователя
func CreateUserAction(ctx context.Context, userAction models.UserAction) (models.UserAction, error) {
	l := logging.GetSugar()

	l.Debug("Выполнение запроса create_user_action")

	db := storage.GetDB()
	sqlStatement := `call create_user_action($1, $2, $3, $4, $5)`
	_, err := db.ExecContext(ctx, sqlStatement, userAction.ID, userAction.Name, userAction.Surname, userAction.Action, userAction.Time.Format(time.RFC3339))

	if err != nil {
		l.Error("Запрос create_user_action выполнен c ошибками: ", err.Error())
		return userAction, err
	}

	l.Debug("Запрос create_user_action выполнен успешно")

	return userAction, nil
}
