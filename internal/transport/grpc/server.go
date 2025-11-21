package transportgrpc

import (
	"fmt"
	"log"
	"net"

	"github.com/enikili/users-service/internal/user"
	userpb "github.com/enikili/users-service/proto/user" // ТОТ ЖЕ ПУТЬ!
	"google.golang.org/grpc"
)

func RunGRPC(svc user.Service) error {
	// Создаем listener на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Создаем gRPC сервер
	grpcServer := grpc.NewServer()

	// Регистрируем наш сервис
	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	log.Printf("🚀 gRPC server started on port 50051")

	// Запускаем сервер
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}