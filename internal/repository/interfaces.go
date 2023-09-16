package repository

import "github.com/marvelalexius/gymbroapp/internal/model"

type (
	IUserRepo interface {
		Store(user *model.User) (*model.User, error)
		Update(user model.User, userID uint) (*model.User, error)
		FindById(id uint) (*model.User, error)
		FindByEmail(email string) (*model.User, error)
	}
)
