package services

import (
	"crowdfunding/dto"
	"crowdfunding/model"
	"crowdfunding/repository"

	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	CreateUser(input dto.UserRequest) (model.User, error)
}

type service struct {
	repository repository.RepositoryUser
}

func ServicesUser(repository repository.RepositoryUser) *service {
	return &service{repository}
}

func (s *service) CreateUser(input dto.UserRequest) (model.User, error) {
	user := model.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = input.Role

	createUser, err := s.repository.RepoCreateUser(user)
	if err != nil {
		return createUser, err
	}

	return createUser, nil

}
