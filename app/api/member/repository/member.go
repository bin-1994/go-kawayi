package repository

import (
	"go-biu/app/api/member/model"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"gorm.io/gorm"
	"time"
)

type MemberDao struct {

}

// 查询昵称是否可用
func (this *MemberDao) CheckNickname(db *gorm.DB, nickname string) *model.User {
	user := &model.User{}
	db.Select("nickname").Where("nickname = ?", nickname).First(&user)
	return user
}

func (this *MemberDao) CheckEmail(db *gorm.DB, email string) *model.User {
	user := &model.User{}
	db.Select("email").Where("nickname = ?", email).First(&user)
	return user
}

func (this *MemberDao) Get(db *gorm.DB, email string) (*model.User) {
	user := &model.User{}
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}
	return user
}

func (this *MemberDao) UpdateLogininfo(db *gorm.DB, ip string, uid int64) {
	db.Exec("UPDATE user SET last_login_ip = ?, last_login_time = ? WHERE uid = ?", ip, time.Now().Unix(), uid)
}

func (this *MemberDao) Create(db *gorm.DB, user *model.User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (this *MemberDao) AddUser(nickname, email, password, ip string, time int64) bool {
	if result := orm.Db.Exec("INSERT INTO `user` (nickname, email, password, reg_ip, create_time) VALUES (?, ?, ?, ?, ?)", nickname, email, utils.UserMd5(password), ip, time); result.Error != nil {
		return true
	}
	return false
}