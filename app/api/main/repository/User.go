package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByEmail(string) (model.UserModel, error)
	FindByNickname(string) (model.UserModel, error)
	FindById(int64) (model.UserModel, error)
	Create(model.UserModel) (model.UserModel, error)
	FindFollow(int64, int) (model.UserProgress, error)
	FindReview(int64, int) (model.UserReview, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{Db: db}
}

func (repository *UserRepository) Create(user model.UserModel) (model.UserModel, error) {
	if result := repository.Db.Table("user").Create(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (repository *UserRepository) FindById(uid int64) (model.UserModel, error) {
	var user model.UserModel
	result := repository.Db.Table("user").Where("uid = ?", uid).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repository *UserRepository) FindByEmail(email string) (model.UserModel, error) {
	var user model.UserModel
	result := repository.Db.Table("user").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repository *UserRepository) FindByNickname(nickname string) (model.UserModel, error) {
	var user model.UserModel
	result := repository.Db.Table("user").Where("nickname = ?", nickname).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repository *UserRepository) FindFollow(uid int64, id int) (model.UserProgress, error) {
	var follow model.UserProgress
	result := repository.Db.Table("follow").Where("uid = ? AND media_id = ?", uid, id).First(&follow)
	if result.Error != nil {
		return follow, result.Error
	}
	return follow, nil
}

func (repository *UserRepository) FindReview(uid int64, id int) (model.UserReview, error) {
	var review model.UserReview
	result := repository.Db.Table("review").Where("uid = ? AND media_id = ?", uid, id).First(&review)
	if result.Error != nil {
		return review, result.Error
	}
	return review, nil
}
