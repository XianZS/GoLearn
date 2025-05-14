package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

const userKey = "session_id"

func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userKey))
	return sessions.Sessions("mySession", store)
}

// AuthSession : UserSession 中间件，用于验证用户是否登录
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get(userKey)
		if sessionID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "请先登录",
				"code":    401,
			})
			return
		}
		c.Next()
	}
}

// SaveSession for User
func SaveSession(c *gin.Context, userID int) {
	session := sessions.Default(c)
	session.Set(userKey, userID)
	err := session.Save()
	if err != nil {
		return
	}
}

// ClearSession for User
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		return
	}
}

// GetSessionID for User
func GetSessionID(c *gin.Context) int {
	session := sessions.Default(c)
	userID := session.Get(userKey)
	if userID == nil {
		return 0
	}
	return userID.(int)
}

// CheckSession for User
func CheckSession(c *gin.Context) bool {
	session := sessions.Default(c)
	userID := session.Get(userKey)
	return userID != nil
}
