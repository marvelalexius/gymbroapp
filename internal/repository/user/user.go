package user

import (
	"github.com/marvelalexius/gymbroapp/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Store(user *model.User) (*model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Update(user model.User, userID uint) (*model.User, error) {
	err := u.db.Model(&user).Where("id = ?", userID).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) FindById(id uint) (*model.User, error) {
	var user *model.User
	err := u.db.Model(&model.User{}).Preload("Gallery").First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *UserRepo) FindByEmail(email string) (*model.User, error) {
	var user *model.User
	err := u.db.Model(&model.User{}).Where("email = ?", email).Preload("Gallery").First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}
