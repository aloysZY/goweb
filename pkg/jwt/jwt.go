package jwt

import (
	"errors"
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
	UserID uint64 `json:"user_id"`
	// UserName             string `json:"user_name"`
	jwt.RegisteredClaims // 内嵌标准的声明
}

// TokenExpireDuration token 有效时间 直接使用配置文件
// const TokenExpireDuration = conf.Config.Jwt.ExpiresTime

// GenToken 生成JWT
func GenToken(userId uint64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	claims := &CustomClaims{
		userId, // 自定义字段
		// username,
		jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.Config.Jwt.ExpiresTime)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.Config.Jwt.ExpiresTime * time.Hour)),
			Issuer:    conf.Config.Jwt.Issuer, // 签发人  随便写
		},
	}
	// 使用指定的签名方法创建签名对象,返回atoken
	// 使用指定的secret签名并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(customSecret)

	// 生成 rToken这个是刷新用的,在这里写每次执行刷新 token 的时候会调用，两个 token 都从新刷新一下
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		// jwt.StandardClaims早期的不建议使用
		// ExpiresAt: time.Now().Add(conf.Config.Jwt.BufferTime * time.Minute).Unix(), // 过期时间
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.Config.Jwt.BufferTime * time.Minute)),
		Issuer:    conf.Config.Jwt.Issuer, // 签发人  随便写
	}).SignedString(customSecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *CustomClaims, err error) {
	// 解析token
	// var token *jwt.Token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	claims = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
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
	// fmt.Printf("jwt userid: %v\n", claims.UserID)
	// fmt.Printf("jwt token: %v\n", token)
	// fmt.Printf("jwt claims: %v\n", claims)
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	// 从旧access token中解析出claims数据
	var claims CustomClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// // 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	// 这里修改成没过期只要是正确的 token 都创建新的 token,这在中间件进行判断

	return
}
