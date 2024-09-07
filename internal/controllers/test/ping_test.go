package test

import (
	"cyclopropane/internal/controllers/api"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 每个接口的对应单点测试应编写在 /internal/controllers/test/xxxx_test.go 中

func TestPing(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", api.Ping)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := gin.H{
		"message": "ping successfully",
	}
	var actualBody gin.H
	_ = json.Unmarshal(w.Body.Bytes(), &actualBody)
	assert.Equal(t, expectedBody, actualBody)
}
