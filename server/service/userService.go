package service

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"linktree_core/modules/db/model"
	"linktree_core/server/dao"
)

var UserService = userService{}

// userService 业务层
type userService struct {
}

func (s userService) GetUser(queryMap map[string][]string) (error, []model.User) {
	if len(queryMap) == 0 {
		return dao.User.DB().All()
	}
	id := queryMap["id"]
	return dao.User.DB().Find(model.User{UserID: uuid.FromStringOrNil(id[0])})
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
