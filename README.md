# Инструкция по запуску и использованию сервиса
Это инструкция по запуску и использованию вашего сервиса, который обогащает данные о людях, включая их возраст,

пол и национальность, и сохраняет их в базе данных PostgreSQ

## Запуск сервиса с использованием Docker
* Убедитесь, что у вас установлен Docker и Docker Compose на вашем компьютере.
* Склонируйте репозиторий с вашим приложением:

```bash
git clone https://github.com/zatrasz75/task_junior.git
```

* Перейдите в каталог с приложением:

```bash
cd task_junior
```

* Запустите приложение с помощью Docker Compose:

```bash
docker-compose up -d
```

* Ваш сервис будет доступен по следующему URL:

```bash
http://localhost:8787
```

* Иначе перенастройте файл .env под себя

```bash
# server
APP_HOST: "localhost"
APP_PORT: "8787"
READ_TIMEOUT: "3s"
WRITE_TIMEOUT : "3s"
IDLE_TIMEOUT: "6s"
SHUTDOWN_TIMEOUT: "10s"

# postgres
HOST_DB: "postgres"
USER_DB: "postgres"
PASSWORD_DB: "postgrespw"
URL_DB: "postgres"
NAME_DB: "Account"
PORT_DB: "5432"
```

* запуск
```bash
go run cmd main.go
```

## REST API методы
### Добавление новых данных о человеке

*  POST-запрос на /api/data

```bash
{
"name": "Dmitriy",
"surname": "Ushakov",
"patronymic": "Vasilevich" // Необязательное поле
}

```
### Получение данных с различными фильтрами и пагинацией
* GET-запрос на /data
*
* GET-запрос на /data?gender=male&age=49&page=1&pageSize=3

### Удаление данных по идентификатору
* DELETE-запрос на /data/{id}

### Изменение данных сущности
* PUT-запрос на /data/{id}

### Частичное изменение данных сущности
*  PATCH-запрос на /data/{id}

### Для добавления новых людей
* POST-запрос на /data

## Логирование
* Ваше приложение включает логирование с уровнями DEBUG и INFO.
* Вы можете найти логи в файле app.log, который настроен в вашем приложении.