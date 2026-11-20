package service

import (
	"errors"
	"fmt"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/repository"
	"go-biu/utils"
	"regexp"
	"time"
	"unicode/utf8"
)

type IUserService interface {
	Login(string, string) (model.UserModel, error)
	Register(string, string, string, string, string) (model.UserModel, error)
	FindById(int64) (model.UserModel, error)
	FindState(int64, int) (model.UserState, error)
	CheckNickname(string) error
	CheckNinameFormat(string) bool
}

type UserService struct {
	Repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{Repository: repository}
}

func (service *UserService) Login(email, password string) (model.UserModel, error) {

	user, _ := service.Repository.FindByEmail(email)

	if user.Uid == 0 {
		return user, errors.New("用户不存")
	}

	if user.Password != utils.UserMd5(password) {
		return user, errors.New("密码错误")
	}

	if user.State == 0 {
		return user, errors.New("该用户已被永久封禁")
	}

	return user, nil
}

func (service *UserService) Register(email, password, nickname, regCode, ip string) (model.UserModel, error) {
	codeKey := fmt.Sprintf("kwy_reg_code:%s:%s", ip, regCode)
	var code model.RegCode
	utils.NewRedis().Get(codeKey, &code)
	if regCode != "" || code.Code != regCode {
		return model.UserModel{}, errors.New("注册码失效")
	}

	if len(nickname) == 0 || !service.CheckNinameFormat(nickname) {
		return model.UserModel{}, errors.New("昵称不可包含特殊字符")
	}

	if len(email) == 0 || !service.CheckEmailFormat(email) {
		return model.UserModel{}, errors.New("邮箱格式不正确")
	}

	if len(password) < 6 || len(password) > 16 {
		return model.UserModel{}, errors.New("密码长度为6-16")
	}

	if user, _ := service.Repository.FindByEmail(email); user.Email == email {
		return model.UserModel{}, errors.New("邮箱已被使用")
	}

	if user, _ := service.Repository.FindByNickname(nickname); user.Nickname == nickname {
		return model.UserModel{}, errors.New("昵称已被使用")
	}

	user := model.UserModel{
		Nickname:   nickname,
		Email:      email,
		Password:   utils.UserMd5(password),
		RegIp:      ip,
		CreateTime: time.Now().Unix(),
	}
	user, err := service.Repository.Create(user)
	if err != nil {
		return model.UserModel{}, errors.New("注册失败")
	}
	utils.NewRedis().Del(codeKey)
	return user, nil
}

func (service *UserService) FindById(id int64) (model.UserModel, error) {
	var user model.UserModel
	user, err := service.Repository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *UserService) FindState(uid int64, media_id int) (model.UserState, error) {
	follow, _ := service.Repository.FindFollow(uid, media_id)
	review, _ := service.Repository.FindReview(uid, media_id)
	state := model.UserState{Progress: follow, Review: review}
	return state, nil
}

func (service *UserService) CheckNickname(nickname string) error {
	if utf8.RuneCountInString(nickname) < 2 || utf8.RuneCountInString(nickname) > 8 {
		return errors.New("昵称长度2-8个字符")
	}

	if !service.CheckNinameFormat(nickname) {
		return errors.New("昵称不可包含特殊字符")
	}

	if user, _ := service.Repository.FindByNickname(nickname); user.Nickname == nickname {
		return errors.New("昵称已被使用")
	}

	return nil
}

func (service *UserService) CheckEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func (service *UserService) CheckNinameFormat(nickname string) bool {
	reg := regexp.MustCompile("^[\u4e00-\u9fa5A-Za-z0-9]+$")
	return reg.MatchString(nickname)
}
