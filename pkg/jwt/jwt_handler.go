package jwt

import (
	"cutego/modules/core/api/v1/response"
	"cutego/modules/core/dao"
	"cutego/pkg/cache"
	"cutego/pkg/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JWTAuth 中间件, 检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 放行的请求先放行
		if doSquare(c) {
			return
		}
		authHeader := c.Request.Header.Get(config.AppEnvConfig.Jwt.Header)
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusUnauthorized,
				"msg":    "请求未携带token, 无权限访问",
			})
			c.Abort()
			return
		}
		// 按空格分割
		authHeaderSplit := strings.SplitN(authHeader, " ", 2)
		if !(len(authHeaderSplit) == 2 && authHeaderSplit[0] == config.AppEnvConfig.Jwt.TokenStartWith) {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusUnauthorized,
				"msg":    "请求头中Token格式有误",
			})
			c.Abort()
			return
		}
		// authHeaderSplit[1]是获取到的tokenString, 我们使用之前定义好的解析JWT的函数来解析它
		currentTokenStr := authHeaderSplit[1]
		claims, err := ParseToken(currentTokenStr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusUnauthorized,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		singleTag := config.AppEnvConfig.Login.Single
		if singleTag {
			token, err := dao.RedisDB.GET(claims.UserInfo.UserName)
			if err == nil {
				if !(token == currentTokenStr) {
					c.JSON(http.StatusOK, gin.H{
						"status": http.StatusUnauthorized,
						"msg":    "您的账号已在其他终端登录, 请重新登录",
					})
					c.Abort()
					return
				}
			}
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// 一些常量
var (
	TokenExpired     error = errors.New("授权已过期")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("令牌非法")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

// CuteClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一些字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CuteClaims struct {
	UserInfo response.UserResponse `json:"userInfo"`
	jwt.StandardClaims
}

// CreateToken 生成一个token
func CreateToken(claims CuteClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 定义Secret
	var tokenSecret = []byte(config.AppEnvConfig.Jwt.TokenSecret)
	return token.SignedString(tokenSecret)
}

// CreateUserToken 生成含有用户信息的token
func CreateUserToken(u *response.UserResponse) (string, error) {
	if config.AppEnvConfig.Jwt.TokenExpired == 0 {
		config.AppEnvConfig.Jwt.TokenExpired = 1
	}
	// 定义JWT的过期时间
	tokenExpired := time.Hour * time.Duration(config.AppEnvConfig.Jwt.TokenExpired)
	// 定义Secret
	var tokenSecret = []byte(config.AppEnvConfig.Jwt.TokenSecret)

	// 创建我们自己的声明
	c := CuteClaims{
		UserInfo: *u, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpired).Unix(), // 过期时间
			Issuer:    "tianjun@odboy.cn",                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(tokenSecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*CuteClaims, error) {
	// 定义Secret
	var tokenSecret = []byte(config.AppEnvConfig.Jwt.TokenSecret)
	token, err := jwt.ParseWithClaims(tokenString, &CuteClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token 过期(授权已过期)
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CuteClaims); ok && token.Valid {
		if config.AppEnvConfig.Login.Single {
			tokenData := cache.GetCache(claims.UserInfo.UserName)
			if tokenData == "" {
				return nil, TokenExpired
			}
		}
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func RefreshToken(tokenString string) (string, error) {
	// 定义Secret
	var tokenSecret = []byte(config.AppEnvConfig.Jwt.TokenSecret)
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CuteClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CuteClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return CreateToken(*claims)
	}
	return "", TokenInvalid
}
