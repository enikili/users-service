PROTOS_PATH := proto
OUT_DIR := .

.PHONY: generate deps clean proto-files

# Создает proto файлы если их нет
proto-files:
	@mkdir -p $(PROTOS_PATH)
	@if [ ! -f "$(PROTOS_PATH)/user.proto" ]; then \
		echo "Creating user.proto..."; \
		cat > $(PROTOS_PATH)/user.proto << 'USER_PROTO'; \
syntax = "proto3"; \
package user; \
option go_package = "github.com/your-org/users-service/proto/user"; \
\
message User { \
  uint32 id = 1; \
  string email = 2; \
  string name = 3; \
} \
\
message CreateUserRequest { \
  string email = 1; \
  string name = 2; \
} \
\
message CreateUserResponse { \
  User user = 1; \
} \
\
message GetUserRequest { \
  uint32 id = 1; \
} \
\
message GetUserResponse { \
  User user = 1; \
} \
\
message UpdateUserRequest { \
  uint32 id = 1; \
  string email = 2; \
  string name = 3; \
} \
\
message UpdateUserResponse { \
  User user = 1; \
} \
\
message DeleteUserRequest { \
  uint32 id = 1; \
} \
\
message DeleteUserResponse { \
  bool success = 1; \
} \
\
message ListUsersRequest { \
  uint32 page = 1; \
  uint32 limit = 2; \
} \
\
message ListUsersResponse { \
  repeated User users = 1; \
  uint32 total = 2; \
} \
\
service UserService { \
  rpc CreateUser (CreateUserRequest) returns (CreateUserResponse); \
  rpc GetUser (GetUserRequest) returns (GetUserResponse); \
  rpc UpdateUser (UpdateUserRequest) returns (UpdateUserResponse); \
  rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse); \
  rpc ListUsers (ListUsersRequest) returns (ListUsersResponse); \
} \
USER_PROTO \
	fi

# Генерирует Go код из proto файлов
generate: proto-files
	@echo "Generating protobuf code..."
	protoc \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTOS_PATH)/*.proto
	@echo "✅ Code generation completed"

# Устанавливает зависимости
deps:
	@echo "Installing dependencies..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go mod tidy
	@echo "✅ Dependencies installed"

# Очищает сгенерированные файлы
clean:
	@echo "Cleaning generated files..."
	find . -name "*.pb.go" -delete
	@echo "✅ Clean completed"

# Полная установка
all: deps generate

# Запуск сервера
run:
	go run cmd/server/main.go

# Проверка компиляции
build:
	go build ./...

# Помощь
help:
	@echo "Available commands:"
	@echo "  make all      - Install deps and generate code"
	@echo "  make generate - Generate protobuf code"
	@echo "  make deps     - Install dependencies"
	@echo "  make clean    - Clean generated files"
	@echo "  make run      - Run the server"
	@echo "  make build    - Build the project"
