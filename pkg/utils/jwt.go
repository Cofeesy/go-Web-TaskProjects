package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//用于加密的字符串
var JWTSECRET = []byte("JWT_SECRET")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

/**
 * @Author Cofeesy
 * @Description //自定义Token颁发
 * @Date 23:21 2022/6/23
 * @Param id uint, username string ”将用户的名字输入作为自定义结构体一部分“
 * @return string,error ”“
 **/
func GenToken(id uint, username string) (string, error) {
	//获取当前时间
	nowTime := time.Now()
	//设置token有效时间
	expiredTime := nowTime.Add(24 * time.Minute)
	//自定义结构体，包含jwt.StandardClaims
	claims := Claims{
		Id:       id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),  //token过期时间，为时间戳（Unix()）
			Issuer:    "memorandumProject", //签发人
		},
	}
	//生成一个tokenClaims加密对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用加密字符串进行加密，返回一个字符串类型的token和err
	token, err := tokenClaims.SignedString(JWTSECRET)
	return token, err
}

/**
 * @Author Cofessy
 * @Description //进行token的验证
 * @Date 23:35 2022/6/23
 * @Param token string
 * @return *Claims, error
 **/
func ParseToken(token string) (*Claims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return JWTSECRET, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { // 校验token
		return claims, nil
	}
	return nil, err
}
