package storage

import (
	"database/sql"
	"fmt"
	"log"
	"main/logging"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Экземпляр базы данных
var db *sql.DB

// Инициализация экземпляра базы данных, параметры выбираются из переменных сред
func InitDB() {
	l := logging.GetSugar()

	l.Debug("Инициализация переменной среды")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	l.Debugf("Подключение к БД: Host: %s:%s; DB:%s; User: %s", dbHost, dbPort, dbName, dbUser)

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))

	if err != nil {
		l.Error("Ошибка при открытии подключения к БД: ", err.Error())
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		l.Error("Ошибка при попытке пинга БД", err.Error())
		panic(err.Error())
	}

	l.Info("Успешно подключен к БД")
}

// Получает экземпляр базы данных
func GetDB() *sql.DB {
	return db
}
