package transportgrpc

import (
	"context"
	"log"

	userpb "github.com/enikili/project-protos/proto/user"
	"github.com/enikili/users-service/internal/user"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	log.Printf("Creating user with email: %s", req.GetEmail())

	createdUser, err := h.svc.CreateUser(req.GetEmail(), req.GetName(), "")
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	response := &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(createdUser.ID),
			Email: createdUser.Email,
			Name:  createdUser.Name,
		},
	}

	log.Printf("User created successfully with ID: %d", createdUser.ID)
	return response, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Printf("Getting user with ID: %d", req.GetId())

	user, err := h.svc.GetUserByID(uint(req.GetId()))
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}

	response := &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    uint32(user.ID),
			Email: user.Email,
			Name:  user.Name,
		},
	}

	return response, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	log.Printf("Updating user with ID: %d", req.GetId())

	updatedUser, err := h.svc.UpdateUser(uint(req.GetId()), req.GetEmail(), req.GetName())
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	response := &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(updatedUser.ID),
			Email: updatedUser.Email,
			Name:  updatedUser.Name,
		},
	}

	log.Printf("User updated successfully: %d", updatedUser.ID)
	return response, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	log.Printf("Deleting user with ID: %d", req.GetId())

	err := h.svc.DeleteUser(uint(req.GetId()))
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return nil, err
	}

	response := &userpb.DeleteUserResponse{
		Success: true,
	}

	log.Printf("User deleted successfully: %d", req.GetId())
	return response, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	log.Printf("Listing all users")

	users, err := h.svc.GetAllUsers()
	if err != nil {
		log.Printf("Error listing users: %v", err)
		return nil, err
	}

	log.Printf("Found %d users in database", len(users))

	var userProtos []*userpb.User
	for _, u := range users {
		log.Printf("Processing user: ID=%d, Email=%s, Name=%s", u.ID, u.Email, u.Name)
		userProtos = append(userProtos, &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
			Name:  u.Name,
		})
	}

	response := &userpb.ListUsersResponse{
		Users: userProtos,
	}

	log.Printf("Returning %d users in response", len(userProtos))
	return response, nil
}
