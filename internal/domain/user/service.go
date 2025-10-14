package user

import (
	"fmt"
	"infra-base-go/pkg/util"
	"time"
)

type Service interface {
	CreateUser(userDTO *CreateUserDTO) (string, *util.HttpError)
	GetUserById(id string) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(userDTO *UpdateUserDTO) *util.HttpError
	DeleteUser(id string) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) CreateUser(userDTO *CreateUserDTO) (string, *util.HttpError) {

	// verify if user already exists
	_, err := s.repository.FindByEmail(userDTO.Email)
	if err == nil {
		return "", util.NewHttpError(fmt.Sprintf("User with email %s already exists", userDTO.Email), 400)
	}

	user := User{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
	}

	id, err := s.repository.Create(&user)

	if err != nil {
		return "", util.NewHttpError(err.Error(), 500)
	}

	return id, nil
}

func (s *service) GetUserById(id string) (User, error) {
	return s.repository.Find(id)
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repository.FindAll()
}

func (s *service) UpdateUser(userDTO *UpdateUserDTO) *util.HttpError {
	user := User{
		ID:        userDTO.ID,
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Password:  userDTO.Password,
		Email:     userDTO.Email,
		UpdatedAt: time.Now().UTC().String(),
	}

	if err := s.repository.Update(&user); err != nil {
		return util.NewHttpError(err.Error(), 500)
	}

	return nil
}

func (s *service) DeleteUser(id string) error {
	return s.repository.Delete(id)
}
