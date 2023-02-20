package service

import (
	"github.com/google/uuid"
	"linktree_core/server/modules/jwt"
)

// userService
type userService struct {
}

// Login 用户登录
func (userService) Login() (error, string) {

	j := jwt.NewJWT()
	token, err := j.CreateToken(j.CreateClaims(jwt.BaseClaims{
		ID:         32423,
		UUID:       uuid.New(),
		Username:   "root",
		LoginPlace: "chongqing",
		LoginIp:    "122.432.212.2",
	}))
	if err != nil {
		return err, ""
	}
	return nil, token
}
