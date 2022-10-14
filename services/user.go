package services

import (
	"github.com/adedaramola/golang-auth/datastore"
	"github.com/adedaramola/golang-auth/datastore/models"
)

type UserService struct {
	store *datastore.DB
}

func NewUserService(db *datastore.DB) *UserService {
	return &UserService{store: db}
}

func (u *UserService) CreateNewUser(user *models.User) error {
	query := `insert into users (fullname, email, password) values ($1, $2, $3)`

	_, err := u.store.Exec(query, user.Fullname, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
