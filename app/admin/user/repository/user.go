package repository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// type UserRepoInterface interface {
// 	List(req *query.ListQuery) (users []model.User, err error)
// 	Get(user model.User) (*model.User, error)
// 	Exist(user model.User) *model.User
// 	ExistByUserID (id string) *model.User
// 	ExistByMobile(mobile string) *model.User
// 	Add(user model.User) (*model.User, error)
// 	Edit(user model.User) (bool, error)
// 	Deleted(user model.User) (bool, error)
// }

// func (repo *UserRepository) List(req *query.ListQuery) (users []model.User, err error) {
// 	db := repo.DB
// 	limit, offset := utils.Paging{req.Limit, req.Page}

// 	return user, nil
// }
