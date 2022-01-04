package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SaveSession(c *gin.Context, key string, value interface{}) bool {
	session := sessions.Default(c)
	session.Set(key, value)
	if err := session.Save(); err != nil {
		return false
	}
	return true
}

func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

func RemoveSession(c *gin.Context, key string) {
	session := sessions.Default(c)
	session.Delete(key)
	_ = session.Save()
}
