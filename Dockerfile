# Шаг 1: Сборка бинарного файла
FROM golang:latest AS builder

# Установим рабочую директорию внутри контейнера
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && update-ca-certificates

# Клонируем репозиторий с GitHub
RUN git clone https://github.com/aid219/rabskie-opovesheniya.git .

# Установим зависимостиsda
RUN go mod tidy

# Скомпилируем бинарный файл
# RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# Шаг 2: Сборка минимального конечного образа
FROM alpine:latest

# Установим корневые сертификаты
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache docker-compose
# Скопируем скомпилированный бинарный файл из этапа сборки
COPY --from=builder /app/main /


# Копируем Docker Compose файл
COPY docker-compose.yml /app/
COPY local.yaml .
# Укажем команду для запуска приложения
CMD ["/main"]