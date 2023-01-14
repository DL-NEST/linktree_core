package code

import (
	"fmt"
)

var (
	OK                   = &Co{httpCode: 200, Code: 20000, Message: "OK"}
	ErrParam             = &Co{httpCode: 400, Code: 10003, Message: "参数有误"}
	ErrUpload            = &Co{httpCode: 404, Code: 20114, Message: "文件上传失败"}
	ErrRouter            = &Co{httpCode: 400, Code: 20115, Message: "请求地址非法或不存在"}
	ErrUserNotFound      = &Co{httpCode: 203, Code: 20102, Message: "用户不存在"}
	ErrPasswordIncorrect = &Co{httpCode: 203, Code: 20104, Message: "密码错误"}

	//InternalServerError = &Co{Code: 10001, Message: "Internal server error"}
	//ErrBind             = &Co{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	//ErrSignParam        = &Co{Code: 10004, Message: "签名参数有误"}
	//ErrValidation         = &Co{Code: 20001, Message: "Validation failed."}
	//ErrDatabase           = &Co{Code: 20002, Message: "Database error."}
	//ErrToken              = &Co{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	//ErrInvalidTransaction = &Co{Code: 20004, Message: "invalid transaction."}
	// user errors
	//ErrEncrypt               = &Co{Code: 20101, Message: "密码加密错误"}
	//ErrTokenInvalid          = &Co{Code: 20103, Message: "Token错误"}
	//ErrUserExistBefor        = &Co{Code: 20105, Message: "用户已存在"}
	//ErrUserCreate            = &Co{Code: 20105, Message: "用户创建错误"}
	//ErrSendSMSTooMany        = &Co{Code: 20109, Message: "已超出当日限制，请明天再试"}
	//ErrVerifyCode            = &Co{Code: 20110, Message: "验证码错误"}
	//ErrEmailOrPassword       = &Co{Code: 20111, Message: "邮箱或密码错误"}
	//ErrTwicePasswordNotMatch = &Co{Code: 20112, Message: "两次密码输入不一致"}
	//ErrRegisterFailed        = &Co{Code: 20113, Message: "注册失败"}
	//ErrCreatedUser           = &Co{Code: 20114, Message: "用户创建失败"}
	//ErrAccessRight           = &Co{Code: 20114, Message: "没有访问权限"}
)

// Co ...
type Co struct {
	httpCode int
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
		return co.httpCode, co.Code, co.Message
	}
}
