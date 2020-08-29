package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Body struct {
	Message string `json:"message,omitempty"`
}

func parseBody(response *httptest.ResponseRecorder) (*Body, error) {
	allBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var body Body
	err = json.Unmarshal(allBody, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func stopIfFailed(t *testing.T) {
	if t.Failed() {
		t.FailNow()
	}
}

func TestHandlePing(t *testing.T) {
	test := assert.New(t)

	ginEngine := gin.Default()
	ginEngine.GET("/ping", HandlePing()...)

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/ping", nil)
	ginEngine.ServeHTTP(response, request)

	test.Nil(err)
	stopIfFailed(t)
	test.Equal(response.Code, 200)

	body, err := parseBody(response)
	test.Nil(err)
	stopIfFailed(t)
	test.Equal(body.Message, "pong")
}
