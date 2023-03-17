package service

import (
	"errors"
	"linktree_core/server/dao"
	"linktree_core/server/model/dto/request"
	"linktree_core/server/model/entity"
	"linktree_core/server/modules/jwt"
	"linktree_core/utils/crypto"
)

// userService
type userService struct {
}

// Login 用户登录
func (userService) Login(request request.LoginRequest) (error, entity.User, string) {

	err, user := dao.User.FindInUserName(request.UserName)
	if err != nil {
		return errors.New("用户不存在"), user, ""
	}
	// 使用md5
	if crypto.Md5V1(user.Password) != request.Password {
		return errors.New("密码错误"), user, ""
	}

	// 签发token
	j := jwt.NewJWT()
	token, err := j.CreateToken(j.CreateClaims(jwt.BaseClaims{
		ID:         user.ID,
		UUID:       user.UUID,
		Username:   user.UserName,
		LoginPlace: request.LoginPlace,
		LoginIp:    request.LoginIp,
	}))
	if err != nil {
		return err, user, ""
	}
	return nil, user, token

}
