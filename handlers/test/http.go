package test

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"server-demo/constants"
	"server-demo/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// HandleURL handle /test/url
//
// @Summary test url
// @Description test url
// @Produce  json
// @Success 200 {string} string	"{"uri":"/prefix/test/url","query": {}}"
// @Router /test/url [GET]
func HandleURL(ctx *gin.Context) {
	request := ctx.Request
	ctx.JSON(http.StatusOK, gin.H{
		"query": request.URL.Query(),
		"uri":   request.URL.RequestURI(),
	})
}

// HandlerHeader handle /test/headers
func HandlerHeader(ctx *gin.Context) {
	request := ctx.Request

	requestHeaders := map[string]interface{}{}
	for k, v := range request.Header {
		if len(v) == 1 {
			requestHeaders[k] = v[0]
		} else {
			requestHeaders[k] = v
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"requestHeaders": requestHeaders,
	})
}

// HandleSleep handle /test/sleep
func HandleSleep(ctx *gin.Context) {
	start := time.Now()
	sleepSecond, err := strconv.Atoi(ctx.Param("second"))
	if err == nil {
		time.Sleep(time.Duration(sleepSecond) * time.Second)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"startTime":  start.Format(constants.TIME_FORMAT_YMDHMS),
		"costTime":   fmt.Sprintf("%d ms", time.Since(start)/time.Millisecond),
		"startStamp": start.Unix(),
	})
}

type cookieParams struct {
	Key      string    `form:"key"`
	Value    string    `form:"value"`
	Path     string    `form:"path"`
	Domain   string    `form:"domain"`
	MaxAge   int       `form:"max-age,default=-1"`
	Expires  time.Time `form:"expires" time_format:"2006-01-02T15:04:05Z07:00"`
	Secure   bool      `form:"secure"`
	HTTPOnly bool      `form:"httponly"`
}

// HandleCookie handle /test/cookie
func HandleCookie(ctx *gin.Context) {
	var params cookieParams
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Println("bind form data error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "bind form data error",
		})
		return
	}

	setCookie := ""
	if len(params.Key) > 0 && len(params.Value) > 0 {
		setCookie += params.Key + "=" + params.Value + "; "

		if params.MaxAge > 0 {
			setCookie += fmt.Sprintf("MaxAge=%d; ", params.MaxAge)
		} else if params.Expires.Unix() > 0 {
			params.MaxAge = int(params.Expires.Unix() - time.Now().Unix())
			setCookie += fmt.Sprintf("Expires=%s; ", params.Expires.Format(constants.TIME_FORMAT_RFC7231))
		}

		if len(params.Path) > 0 {
			setCookie += "Path=" + params.Path + "; "
		} else {
			setCookie += "Path=/; "
		}

		if len(params.Domain) > 0 {
			setCookie += "Domain=" + params.Domain + "; "
		}

		if params.Secure {
			setCookie += "Secure; "
		}

		if params.HTTPOnly {
			setCookie += "HttpOnly; "
		}

		ctx.SetCookie(params.Key, params.Value, params.MaxAge, params.Path, params.Domain, params.Secure, params.HTTPOnly)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"receivedCookie": ctx.Request.Header.Get("cookie"),
		"setCookie":      strings.Trim(strings.TrimSpace(setCookie), ";"),
	})
}

// HandleIP handle /test/ip
func HandleIP(ctx *gin.Context) {
	localIPs, _ := utils.LocalIPs()
	ctx.JSON(http.StatusOK, gin.H{
		"clientIP":  ctx.ClientIP(),
		"serverIPs": localIPs,
	})
}

// HandlePing handle /test/ping
//
// @Summary test ping
// @Description test ping
// @Produce  json
// @Success 200 {string} string	"{"message":"pong"}"
// @Router /test/ping [GET]
func HandlePing() gin.HandlersChain {
	return gin.HandlersChain{
		ping1,
		ping2,
	}
}

func ping1(ctx *gin.Context) {
	fmt.Println("in  func ping1")
	ctx.Next()
	fmt.Println("out func ping1")
}

func ping2(ctx *gin.Context) {
	fmt.Println("	in  func ping2")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	fmt.Println("	out func ping2")
}

type cacheControlParams struct {
	Public         bool      `form:"public"`
	Private        bool      `form:"private"`
	MaxAge         int       `form:"max-age,default=-1"`
	Expires        time.Time `form:"expires" time_format:"2006-01-02T15:04:05Z07:00"`
	NoCache        bool      `form:"no-cache"`
	NoStore        bool      `form:"no-store"`
	MustRevalidate bool      `form:"must-revalidate"`
	LastModified   time.Time `form:"last-modified" time_format:"2006-01-02T15:04:05Z07:00"`
}

// HandleCacheControl handle /test/cache-control
func HandleCacheControl(ctx *gin.Context) {
	request := ctx.Request
	requestCacheControl := request.Header.Get("cache-control")
	// query := request.URL.Query()

	var params cacheControlParams
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Println("bind form data error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "bind form data error",
		})
		return
	}

	setCacheControl := ""
	setExpiresHeader := ""

	if params.Public {
		setCacheControl += "public, "
	} else if params.Private {
		setCacheControl += "private, "
	}

	if params.NoStore {
		setCacheControl += "no-store, "
	}

	if params.NoCache {
		setCacheControl += "no-cache, "
	}

	if params.MustRevalidate {
		setCacheControl += "must-revalidate, "
	}

	if params.MaxAge >= 0 {
		setCacheControl += fmt.Sprintf("max-age=%d, ", params.MaxAge)
	} else if params.Expires.Unix() > 0 {
		setExpiresHeader = params.Expires.Format(constants.TIME_FORMAT_RFC7231)
		ctx.Header("Expires", setExpiresHeader)
	}

	setCacheControl = strings.Trim(strings.Trim(setCacheControl, " "), ",")
	ctx.Header("Cache-Control", setCacheControl)
	ctx.JSON(http.StatusOK, gin.H{
		"Request-Cache-Control":  requestCacheControl,
		"Response-Cache-Control": setCacheControl,
		"Response-Expires":       setExpiresHeader,
	})
}

// HandleMatch handle /test/match
func HandleMatch(ctx *gin.Context) {
	request := ctx.Request
	ifNoneMatch := request.Header.Get("If-None-Match")
	ifModifiedSince := request.Header.Get("If-Modified-Since")
	ifModifiedSinceTime, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", ifModifiedSince)

	var params cacheControlParams
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Println("bind form data error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "bind form data error",
		})
		return
	}

	setLastModified := ""
	if !params.LastModified.IsZero() {
		setLastModified = params.LastModified.Format(constants.TIME_FORMAT_RFC7231)
		ctx.Header("Last-Modified", setLastModified)
	}

	setCacheControl := ""
	if params.MaxAge >= 0 {
		setCacheControl += fmt.Sprintf("max-age=%d", params.MaxAge)
		ctx.Header("Cache-Control", setCacheControl)
	}

	eTag := "\"matching-etag\""
	ctx.Header("ETag", eTag)

	if ifNoneMatch == eTag {
		ctx.JSON(http.StatusNotModified, gin.H{})
	} else if ifModifiedSinceTime.After(params.LastModified) && !params.LastModified.IsZero() {
		ctx.JSON(http.StatusNotModified, gin.H{})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Request-If-None-Match":     ifNoneMatch,
			"Request-If-Modified-Since": ifModifiedSince,
			"Response-Last-Modified":    setLastModified,
			"Response-Cache-Control":    setCacheControl,
			"Response-ETag":             eTag,
		})
	}
}

// ZipOnFly1 handle /zip
// https://stackoverflow.com/a/57434338
func ZipOnFly1(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-type", "application/octet-stream")
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=filename.zip")
	ar := zip.NewWriter(ctx.Writer)

	path1 := "data/harbor-offline-installer-v1.9.2.tgz"
	file1, _ := os.Open(path1)
	f1, _ := ar.Create(path1)
	io.Copy(f1, file1)

	// path2 := "data/vscode.exe"
	// file2, _ := os.Open(path2)
	// f2, _ := ar.Create(path2)
	// io.Copy(f2, file2)

	ar.Close()
}

// ZipOnFly2 ...
// https://stackoverflow.com/a/57434338
func ZipOnFly2(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-type", "application/octet-stream")
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=filename.zip")

	ctx.Stream(func(w io.Writer) bool {
		// Create a zip archive.
		ar := zip.NewWriter(w)

		path1 := "data/mysql.tar"
		file1, _ := os.Open(path1)
		f1, _ := ar.Create(path1)
		io.Copy(f1, file1)

		// path2 := "data/vscode.exe"
		// file2, _ := os.Open(path2)
		// f2, _ := ar.Create(path2)
		// io.Copy(f2, file2)

		ar.Close()
		return false
	})
}

// HandlePostJSON ...
func HandlePostJSON(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	fmt.Println(string(data[:]))
	ctx.Data(http.StatusOK, "application/json", data)
}
