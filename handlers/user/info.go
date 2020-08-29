package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HandleInfo ...
func HandleInfo(ctx *gin.Context) {
	request := ctx.Request

	// get token from header
	xToken := request.Header["X-Token"]
	token := strings.Join(xToken[:], "")
	// fmt.Println("get token from X-Token header is ", token)

	// get token from param
	// query := request.URL.Query()
	// tokenParam, ok := query["token"]
	// if ok {
	// 	fmt.Println("get token from param is", strings.Join(tokenParam[:], ""))
	// }

	// // get token from cookie
	// fmt.Println("get token from cookie ", request.Header["Cookie"])

	if len(token) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    50008,
			"message": "Login failed, unable to get user details.",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 20000,
			"data": map[string]interface{}{
				"introduction": "I am a super administrator",
				"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
				"name":         "Super Admin",
				"roles":        [...]string{"admin"},
			},
		})
	}
}
