package control

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func InitTest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var person Person
		if ctx.ShouldBindQuery(&person) == nil {
			log.Println("====== Only Bind By Query String ======")
			log.Println(person.Name)
			log.Println(person.Address)
		}
		log.Printf("%v", ctx.Param("uri")[1:])
		ctx.String(200, "sad")
	}
}
