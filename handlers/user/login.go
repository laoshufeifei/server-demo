package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FormData is for bind param
type FormData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// HandleLogin ...
func HandleLogin(ctx *gin.Context) {
	var form FormData
	if err := ctx.ShouldBind(&form); err != nil {
		fmt.Println("bind form data error")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    60204,
			"message": "Account and password are incorrect.",
		})
		return
	}

	userName := form.Username
	password := form.Password
	token := "admin-token"

	if len(userName) == 0 || len(password) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    60204,
			"message": "Account and password are incorrect.",
		})
	} else {
		// set cookie is not necessity
		// ctx.SetCookie("vue_admin_template_token", token, 0, "/", "", false, false)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 20000,
			"data": map[string]interface{}{
				"token": token,
			},
		})
	}
}
