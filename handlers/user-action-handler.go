package handlers

import (
	"main/business"
	"main/logging"
	"main/repositories"
	"main/requests"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Обработчик запроса на регистрацию действия пользователя
func UserActionHandler(c echo.Context) error {
	ctx := c.Request().Context()
	l := logging.GetSugar()
	userActionRequest := requests.UserActionRequest{}

	l.Debugf("Привязка модели")
	if err := c.Bind(&userActionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "Ошибка привязки модели").Error())
	}

	l.Debugf("Валидация модели")
	if err := c.Validate(&userActionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "Ошибка валидации").Error())
	}

	l.Debugf("Преобразование модели в бизнес-модель")
	userAction, err := userActionRequest.Map()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	l.Debugf("Выполнение действий над бизнес-моделью")
	business.UpdateUserAction(ctx, userAction)

	l.Debugf("Вызов операции создания действия пользователя")
	updatedUserAction, err := repositories.CreateUserAction(ctx, *userAction)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	l.Debugf("Выполнено успешно")

	return c.JSON(http.StatusCreated, updatedUserAction)
}
