package user

import (
	"github.com/gin-gonic/gin"
)

type requests struct {
	UserName string `json:"userName,omitempty" binding:"required"`
	Password string `json:"password,omitempty"`
}

func (c UserController) Test(ctx *gin.Context) {
	var req requests
	err := c.New(ctx).BuildRequest(&req).Error()
	if err != nil {
		return
	}
	//dao.User.Create(&entity.User{
	//	BaseModel: entity.BaseModel{
	//		ID: 10345,
	//	},
	//	UUID:     uuid.New(),
	//	UserName: "root",
	//	Tel:      13623780271,
	//	Password: "123",
	//	HeadUri:  "/head/3223.jpg",
	//	Role: []string{
	//		"root", "admin",
	//	},
	//})

	c.OK("成功")
}
