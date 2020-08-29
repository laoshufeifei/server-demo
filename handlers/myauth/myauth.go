package myauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicAuth is copy from https://gin-gonic.com/docs/examples/using-basicauth-middleware/
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"abc": "123",
	})
}

// HandleAuthAdmin handle /auth/admin
func HandleAuthAdmin(ctx *gin.Context) {
	user := ctx.MustGet(gin.AuthUserKey).(string)
	ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "secret"})
}
