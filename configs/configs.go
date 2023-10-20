package configs

import (
	"fmt"
	"github.com/zatrasz75/task_junior/pkg/logger"
	"os"
	"time"
)

type (
	Config struct {
		Server   Server
		DataBase DataBase
	}
	Server struct {
		AddrPort     string
		AddrHost     string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
		ShutdownTime time.Duration
	}
	DataBase struct {
		ConnStr string //postgres://postgres:postgrespw@localhost:49153/Account

		Host     string // postgres
		User     string // postgres
		Password string // postgrespw
		Url      string // localhost
		Name     string // Account
		Port     string // 49153
	}
)

func initDB() string {
	c := &Config{
		DataBase: DataBase{
			Host:     os.Getenv("HOST_DB"),
			User:     os.Getenv("USER_DB"),
			Password: os.Getenv("PASSWORD_DB"),
			Url:      os.Getenv("URL_DB"),
			Port:     os.Getenv("PORT_DB"),
			Name:     os.Getenv("NAME_DB"),
		},
	}
	connStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.DataBase.Host, c.DataBase.User, c.DataBase.Password, c.DataBase.Url, c.DataBase.Port, c.DataBase.Name,
	)

	return connStr
}

func New() *Config {
	readTimeoutStr := os.Getenv("READ_TIMEOUT")
	var readTimeout time.Duration
	if readTimeoutStr != "" {
		var err error
		readTimeout, err = time.ParseDuration(readTimeoutStr)
		if err != nil {
			logger.Error("ошибки парсинга времени", err)
		}
	}

	writeTimeoutStr := os.Getenv("WRITE_TIMEOUT")
	var writeTimeout time.Duration
	if writeTimeoutStr != "" {
		var err error
		writeTimeout, err = time.ParseDuration(writeTimeoutStr)
		if err != nil {
			logger.Error("ошибки парсинга времени", err)
		}
	}

	idleTimeoutStr := os.Getenv("IDLE_TIMEOUT")
	var idleTimeout time.Duration
	if idleTimeoutStr != "" {
		var err error
		idleTimeout, err = time.ParseDuration(idleTimeoutStr)
		if err != nil {
			logger.Error("ошибки парсинга времени", err)
		}
	}

	shutdownTimeStr := os.Getenv("SHUTDOWN_TIMEOUT")
	var shutdownTime time.Duration
	if shutdownTimeStr != "SECRET_KEY_TOKEN" {
		var err error
		shutdownTime, err = time.ParseDuration(shutdownTimeStr)
		if err != nil {
			logger.Error("ошибки парсинга времени", err)
		}
	}

	return &Config{
		Server: Server{
			AddrHost:     os.Getenv("APP_HOST"),
			AddrPort:     os.Getenv("APP_PORT"),
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
			ShutdownTime: shutdownTime,
		},
		DataBase: DataBase{
			ConnStr: initDB(),
		},
	}
}
