package gorm

import (
	"github.com/micro-in-cn/starter-kit/srv/account/domain/model"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) FindById(id int64) (*model.User, error) {
	user := model.User{}
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	} else if user.Id == 0 {
		return nil, nil
	} else {
		return &user, nil
	}
}

func (r *userRepository) FindByName(name string) (*model.User, error) {
	user := model.User{}
	if err := db.Where("name = ?", name).First(&user).Error; err == nil {
		return &user, nil
	} else if user.Id == 0 {
		return nil, nil
	} else {
		return &user, nil
	}
}

func (r *userRepository) Add(user *model.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) List(page, size int) ([]*model.User, error) {
	list := make([]*model.User, 0)
	err := db.Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error

	return list, err
}
