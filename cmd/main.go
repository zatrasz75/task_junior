package main

import (
	"github.com/joho/godotenv"
	"github.com/zatrasz75/task_junior/configs"
	"github.com/zatrasz75/task_junior/internal/app"
	"github.com/zatrasz75/task_junior/pkg/logger"
)

// Init Вызывается перед main() и загружает значения из файла .env в систему.
func init() {
	if err := godotenv.Load(); err != nil {
		logger.Error("Файл .env не найден.", err)
	}
}

func main() {
	cfg := configs.New()
	app.Run(cfg)
}
