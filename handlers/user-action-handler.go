package handlers

import (
	"main/repositories"
	"main/requests"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Обработчик запроса на регистрацию действия пользователя
func UserActionHandler(c echo.Context) error {
	userActionRequest := requests.UserActionRequest{}

	//Привязка модели
	if err := c.Bind(&userActionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "Ошибка привязки модели").Error())
	}

	//Валидация модели
	if err := c.Validate(&userActionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "Ошибка валидации").Error())
	}

	//Преобразование модели в бизнес-модель
	userAction, err := userActionRequest.Map()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//Выполнение действий над бизнес-моделью
	userAction.ID += 1

	//Вызов операции создания действия пользователя
	updatedUserAction, err := repositories.CreateUserAction(*userAction)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, updatedUserAction)
}
