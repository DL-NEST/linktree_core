package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"linktree_core/server/dao"
	"linktree_core/server/entity"
)

var UserService = userService{}

// userService 业务层
type userService struct {
}

func (s userService) GetUser(queryMap map[string][]string) (error, []entity.User) {
	if len(queryMap) == 0 {
		return dao.User.DB().All()
	}
	id := queryMap["id"]
	user, err := uuid.Parse(id[0])
	if err != nil {
		return nil, nil
	}
	return dao.User.Finds(&entity.User{UserID: user})
}

// AddUser 注册
func (s userService) AddUser(c *gin.Context) {

}

// UpdateUser 注册
func (s userService) UpdateUser(c *gin.Context) {

}

// DeleteUser 注册
func (s userService) DeleteUser(c *gin.Context) {

}
