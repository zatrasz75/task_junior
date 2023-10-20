package app

import (
	"github.com/zatrasz75/task_junior/configs"
	"github.com/zatrasz75/task_junior/internal/models"
	"github.com/zatrasz75/task_junior/internal/repository"
	"github.com/zatrasz75/task_junior/pkg/api"
	"github.com/zatrasz75/task_junior/pkg/logger"
	"github.com/zatrasz75/task_junior/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *configs.Config) {
	pg, err := postgres.New(cfg.DataBase.ConnStr)
	if err != nil {
		logger.Fatal("нет соединения с postgreSQL", err)
	}

	r := repository.New(pg)
	err = r.PG.AutoMigrate(&models.Person{})
	if err != nil {
		return
	}

	server := api.New(cfg, r)

	go func() {
		err = server.Start()
		if err != nil {
			logger.Fatal("Ошибка при запуске сервера:", err)
		}
	}()

	// Канал для управления остановкой приложений
	serverDoneCh := make(chan os.Signal, 1)
	signal.Notify(serverDoneCh, os.Interrupt, syscall.SIGTERM)
	// Ожидание сигнала завершения работы сервера
	select {
	case s := <-serverDoneCh:
		logger.Info("принят сигнал прерывания %s", s.String())
	}
	err = server.Stop()
	if err != nil {
		logger.Error("не удалось завершить работу сервера", err)
	}
	logger.Info("Сервер успешно выключен")

}
