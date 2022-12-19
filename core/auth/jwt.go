package auth

import (
	"bet/core/key"
	"bet/model"
	"bet/utils"
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type UserInfo struct {
	Account    string
	Uid        int64
	Pid        int64
	Xid        string
	Email      string
	ParentPath string
}

type login struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GinJWTMiddleware() *jwt.GinJWTMiddleware {
	// the auth middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "2xbet zone",
		Key:         []byte("2xbet2xbet...admin888..."),
		Timeout:     time.Hour * 480,
		MaxRefresh:  time.Hour * 480,
		IdentityKey: "Uid",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserInfo); ok {
				return jwt.MapClaims{
					"Uid":        v.Uid,
					"Account":    v.Account,
					"Email":      v.Email,
					"Pid":        v.Pid,
					"ParentPath": v.ParentPath,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			uid := claims["Uid"].(float64)
			v, ok := claims["Pid"]
			var pid int64
			if ok {
				a := v.(float64)
				pid = int64(a)
			}
			return &UserInfo{
				Uid:        int64(uid),
				Account:    claims["Account"].(string),
				Email:      claims["Email"].(string),
				Pid:        pid,
				ParentPath: claims["ParentPath"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", errors.New("incorrect Email or Password")
			}
			userID := loginVals.Account
			password := utils.Md5(utils.Md5Key + loginVals.Password)
			var user model.UserBase
			utils.DB.First(&user, "account=? and password=? and status=1", userID, password)
			if user.Uid == 0 {
				return "", errors.New("incorrect Email or Password")
			}
			u := &UserInfo{
				Account:    userID,
				Uid:        user.Uid,
				Email:      user.Email,
				Xid:        user.Xid,
				Pid:        user.Pid,
				ParentPath: user.ParentPath,
			}
			err := utils.RedisSet(fmt.Sprintf(key.RK_JWT_USERINFO_UID, user.Uid), u, time.Hour*480)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			return u, nil

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {

			v, ok := data.(*UserInfo)
			if !ok {
				return false
			}
			err := utils.RedisGet(fmt.Sprintf(key.RK_JWT_USERINFO_UID, v.Uid), v)
			if err != nil {
				return false
			}

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: auth",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "2xbet",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}
func GetUserInfo(gx *gin.Context) *UserInfo {
	uid := gx.Keys["Uid"]
	info := uid.(*UserInfo)
	return info
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
