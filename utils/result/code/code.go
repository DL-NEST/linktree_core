package code

import (
	"fmt"
)

var (
	OK          = &Co{HttpCode: 200, Code: 0, Message: "OK"}
	OKorCreated = &Co{HttpCode: 201, Code: 0, Message: "资源上传成功"}

	ErrParam       = &Co{HttpCode: 400, Code: 2001, Message: "参数有误"}
	ErrAuth        = &Co{HttpCode: 401, Code: 2002, Message: "没有访问权限"}
	ErrAccessRight = &Co{HttpCode: 403, Code: 2003, Message: "没有访问资源的权限"}
	ErrNotFound    = &Co{HttpCode: 404, Code: 2004, Message: "请求地址和资源非法或不存在"}
	ErrNotAllowed  = &Co{HttpCode: 405, Code: 2005, Message: "资源修改失败"}
	ErrUpload      = &Co{HttpCode: 405, Code: 2005, Message: "资源上传失败"}
	ErrProxyAuth   = &Co{HttpCode: 407, Code: 2006, Message: "第三方授权失败"}
	ErrServeState  = &Co{HttpCode: 409, Code: 2007, Message: "服务正在维护升级"}
	ErrOldAPI      = &Co{HttpCode: 410, Code: 2008, Message: "接口已经弃用"}
	ErrServe       = &Co{HttpCode: 500, Code: 2009, Message: "服务器不知所错"}
)

// Co ...
type Co struct {
	HttpCode int
	Code     int
	Message  string
}

func (code *Co) LogError() string {
	return fmt.Sprintf("Code :%d | Message :%v", code.Code, code.Message)
}

func Decode(co *Co) (int, int, string) {
	if co == nil {
		return OK.Code, OK.Code, OK.Message
	} else {
		return co.HttpCode, co.Code, co.Message
	}
}
