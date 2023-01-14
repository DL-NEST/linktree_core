package service

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"linktree_server/modules/db"
	"linktree_server/modules/db/dao"
	"linktree_server/modules/db/model"
)

var UserService = userService{}

// userService 业务层
type userService struct {
}

func (s userService) GetUser(queryMap map[string][]string) []model.User {
	if len(queryMap) == 0 {
		return dao.User.DB(db.DB).All()
	}
	id := queryMap["id"]
	return dao.User.DB(db.DB).Find(model.User{UserID: uuid.FromStringOrNil(id[0])})
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
