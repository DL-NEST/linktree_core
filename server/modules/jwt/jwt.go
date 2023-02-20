package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"linktree_core/utils/encoding"
	"time"
)

type JWT struct {
	SigningKey []byte
}

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID         uint
	UUID       uuid.UUID
	Username   string
	LoginPlace string
	LoginIp    string
}

var (
	TokenExpired     = errors.New("token is expired")           // 令牌已过期
	TokenNotValidYet = errors.New("token not active yet")       // 令牌尚未激活
	TokenMalformed   = errors.New("that's not even a token")    // 那甚至不是令牌
	TokenInvalid     = errors.New("couldn't handle this token") // 无法处理此令牌
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(viper.GetString("server.auth-jwt.signing-key")),
	}
}

// CreateClaims 创建Claims
func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	ep, err := encoding.ParseDuration(viper.GetString("server.auth-jwt.jwt-expires-time"))
	if err != nil {
		ep = 7 * 24 * time.Hour
	}
	claims := CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),         // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间 7天  配置文件
			Issuer:    "linktree",                             // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
