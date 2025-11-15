package repository

import (
	"crowdfunding/model"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	RepoCreateUser(user model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RepoCreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
