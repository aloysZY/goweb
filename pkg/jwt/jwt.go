package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/aloysZy/goweb/global/conf"
	"github.com/golang-jwt/jwt/v4"
)

// customSecret 用于加盐的字符串,随便写,这个现在从配置文件读取存在问题
// panic: runtime error: invalid memory address or nil pointer dereference
// var customSecret = []byte(conf.Config.SigningKey)
var customSecret = []byte("goweb")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return customSecret, nil
}

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               uint64 `json:"user_id"`
	UserName             string `json:"user_name"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// TokenExpireDuration token 有效时间 直接使用配置文件
// const TokenExpireDuration = conf.Config.Jwt.ExpiresTime

// GenToken 生成JWT
func GenToken(userId uint64, username string) (string, error) {
	// 创建一个我们自己的声明
	claims := &CustomClaims{
		userId, // 自定义字段
		username,
		jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.Config.Jwt.ExpiresTime)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.Config.Jwt.ExpiresTime * time.Hour)),
			Issuer:    conf.Config.Jwt.Issuer, // 签发人  随便写
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(customSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *CustomClaims, err error) {
	// 解析token
	var token *jwt.Token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	claims = new(CustomClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	// token, err = jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
	// 	// 直接使用标准的Claim则可以直接使用Parse方法
	// 	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
	// 	return customSecret, nil
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// 对token对象中的Claim进行类型断言
	// if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
	// 	return claims, nil
	// }
	if !token.Valid {
		err = errors.New("invalid token")
	}
	fmt.Printf("jwt userid: %v\n", claims.UserID)
	fmt.Printf("jwt token: %v\n", token)
	fmt.Printf("jwt claims: %v\n", claims)
	return
}
