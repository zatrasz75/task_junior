package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/zatrasz75/task_junior/configs"
	"github.com/zatrasz75/task_junior/internal/controllers"
	"github.com/zatrasz75/task_junior/internal/repository"
	"github.com/zatrasz75/task_junior/pkg/logger"
	"net/http"
	"time"
)

// API представляет собой приложение с набором обработчиков.
type API struct {
	r            *mux.Router
	port         string
	host         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	shutdownTime time.Duration
	srv          *http.Server
	controllers  *controllers.Server
}

// Router возвращает маршрутизатор запросов.
func (api *API) Router() *mux.Router {
	return api.r
}

// New Конструктор
func New(cfg *configs.Config, pg *repository.Store) *API {
	api := &API{
		r:            mux.NewRouter(),
		host:         cfg.Server.AddrHost,
		port:         cfg.Server.AddrPort,
		readTimeout:  cfg.Server.ReadTimeout,
		writeTimeout: cfg.Server.WriteTimeout,
		idleTimeout:  cfg.Server.IdleTimeout,
		shutdownTime: cfg.Server.ShutdownTime,
		controllers:  &controllers.Server{PG: pg},
	}

	// Регистрируем обработчики API.
	api.endpoints()

	return api
}

// Start Метод для запуска сервера
func (api *API) Start() error {
	api.srv = &http.Server{
		Addr:         api.host + ":" + api.port,
		Handler:      api.r,
		ReadTimeout:  api.readTimeout,
		WriteTimeout: api.writeTimeout,
		IdleTimeout:  api.idleTimeout,
	}
	logger.Info("Запуск сервера на http://" + api.srv.Addr)

	go func() {
		err := api.srv.ListenAndServe()
		if err != nil {
			logger.Error("Остановка сервера", err)
			return
		}
	}()

	return nil
}

// Stop останавливает сервер, используя метод Shutdown.
func (api *API) Stop() error {
	if api.srv != nil {
		// Создаем контекст с таймаутом для остановки сервера.
		ctx, cancel := context.WithTimeout(context.Background(), api.shutdownTime)
		defer cancel()

		// Вызываем метод Shutdown для остановки сервера.
		if err := api.srv.Shutdown(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (api *API) endpoints() {
	api.r.HandleFunc("/api/data", api.controllers.ReceiveSave).Methods(http.MethodPost)
	api.r.HandleFunc("/data", api.controllers.GetData).Methods(http.MethodGet)
	api.r.HandleFunc("/data/{id}", api.controllers.DeleteData).Methods(http.MethodDelete)
	api.r.HandleFunc("/data/{id}", api.controllers.UpdateData).Methods(http.MethodPut)
	api.r.HandleFunc("/data/{id}", api.controllers.PartialUpdateData).Methods(http.MethodPatch)
}
