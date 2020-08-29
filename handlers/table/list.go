package table

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("table.init")
	rand.Seed(time.Now().UnixNano())
}

type tableInfo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Author      string    `json:"author"`
	DisplayTime time.Time `json:"display_time"`
	PageViews   int       `json:"pageviews"`
}

var statusMap = []string{
	"published",
	"draft",
	"deleted",
}

// https://blog.csdn.net/oqqYuan1234567890/article/details/89287717
func (d tableInfo) MarshalJSON() ([]byte, error) {
	type Alias tableInfo
	return json.Marshal(&struct {
		Alias
		DisplayTime string `json:"display_time"`
	}{
		Alias:       (Alias)(d),
		DisplayTime: d.DisplayTime.Format("2006-01-02 15:04:05"),
	})
}

// HandleTableList ...
func HandleTableList(ctx *gin.Context) {
	maxTable := 30
	items := make([]tableInfo, 30)
	// fmt.Println(items)
	for i := 0; i < maxTable; i++ {
		items[i].ID = i
		items[i].Title = fmt.Sprintf("tile-%d", i)
		items[i].Status = statusMap[rand.Intn(3)]
		items[i].Author = "author-name"
		items[i].DisplayTime = time.Now().Add(time.Duration(-1*rand.Intn(10000)) * time.Second)
		items[i].PageViews = rand.Intn(1000)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": len(items),
			"items": items,
		},
	})
}
