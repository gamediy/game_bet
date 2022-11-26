package core

import (
	"bet/core/key"
	"bet/model"
	"bet/utils"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type UserInfo struct {
	Account string
	Uid     int32
	Xid     string
	Email   string
}

type login struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GinJWTMiddleware() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "2xbet zone",
		Key:         []byte("2xbet2xbet...admin888..."),
		Timeout:     time.Hour * 48,
		MaxRefresh:  time.Hour * 48,
		IdentityKey: "Uid",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserInfo); ok {
				return jwt.MapClaims{
					"Uid":     v.Uid,
					"Account": v.Account,
					"Email":   v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserInfo{
				Uid:     claims["Uid"].(int32),
				Account: claims["Account"].(string),
				Email:   claims["Email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Account
			password := utils.Md5(utils.Md5Key + loginVals.Password)
			var user model.UserBase
			utils.DbMain.First(&user, "account=? and password=? and status=1", userID, password)
			if user.Uid == 0 {
				return nil, jwt.ErrFailedAuthentication
			}
			u := &UserInfo{
				Account: userID,
				Uid:     user.Uid,
				Email:   user.Email,
				Xid:     user.Xid,
			}
			err := utils.RedisSet(fmt.Sprintf(key.RK_JWT_USERINFO_UID, user.Uid), u, time.Hour*48)
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
			utils.RedisGet(fmt.Sprintf(key.RK_JWT_USERINFO_UID, v.Uid), v)
			if v == nil {
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
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
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
