# Используем официальный образ Go
FROM golang:1.20.6

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod .

# Загружаем зависимости
RUN go mod download

# Копируем код
COPY .. .

# Сборка приложения
RUN go build -o main

# Запуск приложения
CMD ["./main"]