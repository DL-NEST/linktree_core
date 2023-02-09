package service

import (
	"errors"
	"linktree_core/modules/database/db/model"
	"linktree_core/utils/gos"
)

type UserRegisterInfo struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Time     string `form:"time" json:"time" binding:"required,min=8,max=40"`
}

// UserLoginInfo 用户登录后的输出输出
type UserLoginInfo struct {
	UserInfo model.User
	Token    string `form:"token" json:"token"`
}

func JudgeToken(username string, token string) bool {

	return true
}

//func Login(userName string, passWord string) (bool, *code.Co, UserLoginInfo) {
//	//boo, msg, user := sqls.FindUserInNameOrPassword(userName, passWord)
//	//return boo, msg, UserLoginInfo{
//	//	UserInfo: user,
//	//	Token:    "",
//	//}
//}

// Register 注册
//func Register(param control.UserRegisterParam) string {
//	DB.AddUser(&DB.User{
//		ID:         uuid.NewV4(),
//		Name:       param.UserName,
//		Tel:        param.Tel,
//		CreateTime: time.Now(),
//	})
//	return "ok"
//}

func FirstLogin(username string, password string) error {
	err, usn, pwd := gos.ReadPassFile()
	if err != nil {
		return err
	}
	if username == usn && password == pwd {
		return nil
	} else {
		return errors.New("the username or password is incorrect")
	}
}