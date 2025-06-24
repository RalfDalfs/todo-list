# Stage 1: Сборка
FROM golang:1.24.3 as builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /app/app

# Stage 2: Финальный образ
FROM ubuntu:alpine
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app

# Копируем бинарник из первого этапа
COPY --from=builder /app/app /app/app
# Копируем веб-файлы из локального контекста
COPY web /app/web

ENV PORT 7450
EXPOSE $PORT
CMD ["./app"]