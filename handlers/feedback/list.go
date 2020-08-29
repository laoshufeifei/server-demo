package feedback

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type feedbackInfo struct {
	ID        int       `json:"id"`
	User      string    `json:"user"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	Contact   int       `json:"contact"`
}

var feedTypes = []string{
	"poi",
	"routing",
	"camera",
	"other",
}

type feedbackParams struct {
	Page        int       `form:"page,default=1"`
	Limit       int       `form:"limit,default=30"`
	CreatedFrom time.Time `form:"createdFrom" time_format:"2006-01-02T15:04:05Z07:00"` // or time_format:"unix"
	CreatedTo   time.Time `form:"createdTo"   time_format:"2006-01-02T15:04:05Z07:00"`
	SortKey     string    `form:"sortKey,default=id"`
	Decrease    bool      `form:"decrease"`
	Type        string    `form:"type"`
}

// https://blog.csdn.net/oqqYuan1234567890/article/details/89287717
func (d feedbackInfo) MarshalJSON() ([]byte, error) {
	type Alias feedbackInfo
	return json.Marshal(&struct {
		Alias
		CreatedTime string `json:"createdAt"`
	}{
		Alias:       (Alias)(d),
		CreatedTime: d.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// HandleFeedbackList ...
func HandleFeedbackList(ctx *gin.Context) {
	var params feedbackParams
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Println("bind form data error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    60204,
			"message": "bind form data error",
		})
		return
	}

	fromStamp := params.CreatedFrom.Unix()
	toStamp := params.CreatedTo.Unix()
	sortByID := strings.Index(params.SortKey, "id") != -1
	sortByCreatedAt := strings.Index(params.SortKey, "createdAt") != -1
	decrease := params.Decrease

	limit := params.Limit
	items := make([]feedbackInfo, limit)
	for i := 0; i < limit; i++ {
		var feedType string
		if len(params.Type) > 0 {
			feedType = params.Type
		} else {
			feedType = feedTypes[rand.Intn(4)]
		}

		var randTime time.Time
		if fromStamp > 0 && toStamp > 0 {
			stamp := rand.Int63n(toStamp-fromStamp) + fromStamp
			randTime = time.Unix(stamp, 0)
		} else {
			randTime = time.Now().Add(time.Duration(-1*rand.Intn(9999)) * time.Minute)
		}

		items[i].ID = i + 1
		items[i].Content = fmt.Sprintf("Content-%d xxxxxx", i+1)
		items[i].Type = feedType
		items[i].User = "feedback-user-name"
		items[i].CreatedAt = randTime
		items[i].Contact = rand.Intn(1000)
	}

	if sortByCreatedAt {
		if decrease {
			sort.Slice(items, func(i, j int) bool {
				return items[i].CreatedAt.After(items[j].CreatedAt)
			})
		} else {
			sort.Slice(items, func(i, j int) bool {
				return items[i].CreatedAt.Before(items[j].CreatedAt)
			})
		}
	} else if sortByID {
		if decrease {
			sort.Slice(items, func(i, j int) bool {
				return items[i].ID > items[j].ID
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"total": len(items),
			"items": items,
		},
	})
}
