package storage

import (
	"github.com/Studio56School/university/internal/model"
)

type IRepository interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}
