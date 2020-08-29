package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleLogout ...
func HandleLogout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data": "success",
	})
}
