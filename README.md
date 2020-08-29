README
=======


运行
----

```console
go env -w GOPROXY="https://goproxy.cn,direct"
go env -w GO111MODULE="on"
go run main.go
```


API
----

/test/ping

/test/ip
/test/url
/test/header

/test/time
/test/sleep/{second}

/test/cookie
/test/cookie?key=a&value=1&domain=127.0.0.1&path=/test&httponly=true&max-age=10
/test/cookie?key=a&value=1&expires=2020-02-16T07:04:05Z
/test/cookie?key=a&value=1&path=%20 (space)

/test/cache-control?max-age=1000&no-cache=1&no-store=true&must-revalidate=1
/test/cache-control?expires=2020-02-16T07:04:05Z&public=true&private=false

/test/match
If-None-Match: "\"matching-etag\""

/test/match?last-modified=2020-01-16T14:35:24Z
If-Modified-Since: Thu, 16 Jan 2020 15:35:24 GMT

/test/match
max-age=100
Etag=abc

If-None-Match="abc" 304


/test/mysql/ping
/test/redis/ping


/user/login?username=xxx&password=xxx
/user/logout
/user/info

/table/list

/feedback/list?page=1&limit=30&type=x&sortKey=xx&descrease=true&createdFrom=2006-01-02T15:04:05Z&createdTo=2006-01-02T15:04:05Z

ws
/echo

static file
/data/

upload file
curl -X POST http://127.0.0.1:9090/data -F "file=@./abc.log"  -H "Content-Type: multipart/form-data"

static file with auth
/auth/data

auth
/auth/admin
