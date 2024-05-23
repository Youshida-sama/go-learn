package handlers

import (
	"main/logging"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Обработчик пустых запросов, необходим для заглушек
func BlankHandler(c echo.Context) error {
	l := logging.GetSugar()
	l.Debug("Вызвана заглушка")
	return c.String(http.StatusOK, "Ну и что ты тут забыл?")
}
