FROM golang:latest as builder
LABEL authors="zatrasz@ya.ru"

# Создание рабочий директории
RUN mkdir -p /app

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта внутрь контейнера
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o Junior ./cmd/main.go

# Второй этап: создание production образ
FROM ubuntu AS chemistry

WORKDIR /app

RUN apt-get update

COPY --from=builder /app/Junior ./
COPY ./ ./

CMD ["./Junior"]