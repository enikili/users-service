package user

import (
	"errors"
)

type Service interface {
	CreateUser(email, name, password string) (*User, error)
	GetUserByID(id uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id uint, email, name string) (*User, error)
	DeleteUser(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(email, name, password string) (*User, error) {
	// Проверяем, существует ли пользователь с таким email
	existingUser, _ := s.repo.GetByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	user := &User{
		Email:    email,
		Name:     name,
		Password: password, // В реальном приложении здесь должно быть хеширование
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetUserByEmail(email string) (*User, error) {
	return s.repo.GetByEmail(email)
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateUser(id uint, email, name string) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Если email изменился, проверяем что он уникальный
	if email != user.Email {
		existingUser, _ := s.repo.GetByEmail(email)
		if existingUser != nil {
			return nil, errors.New("user with this email already exists")
		}
		user.Email = email
	}

	user.Name = name

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}