package main

import (
	"log"

	"github.com/enikili/users-service/internal/database"
	"github.com/enikili/users-service/internal/user"
	transportgrpc "github.com/enikili/users-service/internal/transport/grpc"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Создание репозитория и сервиса
	repo := user.NewRepository(database.DB)
	svc := user.NewService(repo)

	// Запуск gRPC сервера
	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}