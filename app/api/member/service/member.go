package service

import (
	"errors"
	"fmt"
	"go-biu/app/api/member/model"
	"go-biu/app/api/member/repository"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

type MemberService struct {
}

func (this *MemberService) Login(email, password, ip string) (string, error) {
	email = strings.TrimSpace(email)
	password = strings.Trim(password, " ")

	if len(email) == 0 {
		return "", errors.New("邮箱不能为空")
	}
	if len(password) == 0 {
		return "", errors.New("密码不能为空")
	}
	if !this.CheckEmailFormat(email) {
		return "", errors.New("邮箱格式不正确")
	}
	memberDao := repository.MemberDao{}
	db := orm.Db
	user := memberDao.Get(db, email)
	if user == nil || user.State != 1 {
		return "", errors.New("用户不存在或被禁用")
	}

	if utils.UserMd5(password) != user.Password {
		return "", errors.New("密码错误")
	}

	// 更新登录ip和时间
	memberDao.UpdateLogininfo(db, ip, user.Uid)

	token := utils.CreateToken("21d6e79b3b2a0f9ad0373441e4f4a03e", &utils.UserClaims{
		Uid:   user.Uid,
		Email: user.Email,
	}, time.Hour*60*7)

	return token, nil
}

func (this *MemberService) Register(email, password, nickname, regCode, ip string) error {

	codeKey := fmt.Sprintf("kwy_reg_code:%s:%s", ip, regCode)
	var code model.RegCode
	utils.NewRedis().Get(codeKey, &code)

	if regCode == "" || code.Code != regCode {
		return errors.New("注册码已失效")
	}

	if len(nickname) == 0 || !this.CheckNinameFormat(nickname) {
		return errors.New("昵称不可包含特殊字符")
	}

	if len(password) < 6 || len(password) > 16 {
		return errors.New("密码长度为6-16")
	}

	if len(email) == 0 || !this.CheckEmailFormat(email) {
		return errors.New("邮箱格式不正确")
	}

	memberDao := repository.MemberDao{}
	db := orm.Db
	if user := memberDao.CheckNickname(db, nickname); user.Username == nickname {
		return errors.New("昵称已被使用")
	}

	if user := memberDao.CheckEmail(db, email); user.Email == email {
		return errors.New("邮箱已被使用")
	}

	user := model.User{
		Nickname:   nickname,
		Email:      email,
		Password:   utils.UserMd5(password),
		RegIp:      ip,
		CreateTime: time.Now().Unix(),
	}
	if err := memberDao.Create(db, &user); err != nil {
		return errors.New("注册失败")
	}

	utils.NewRedis().Del(codeKey)
	return nil
}

func (this *MemberService) CheckNickname(nickname string) error {
	if utf8.RuneCountInString(nickname) < 2 || utf8.RuneCountInString(nickname) > 8 {
		return errors.New("昵称长度2-8个字符")
	}

	if !this.CheckNinameFormat(nickname) {
		return errors.New("昵称不可包含特殊字符")
	}

	if user := new(repository.MemberDao).CheckNickname(orm.Db, nickname); user.Nickname == nickname {
		return errors.New("昵称已被使用")
	}

	return nil
}

func (this *MemberService) CheckEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func (this *MemberService) CheckNinameFormat(nickname string) bool {
	reg := regexp.MustCompile("^[\u4e00-\u9fa5A-Za-z0-9]+$")
	return reg.MatchString(nickname)
}
