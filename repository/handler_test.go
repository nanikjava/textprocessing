package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"rockt/repository/model"
	"testing"
)

func doPost(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestInvalidRequest(t *testing.T) {
	body := "{}"
	d, _ := os.Getwd()
	router := SetupRouter(d)
	w := doPost(router, "POST", "/", body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response model.ResponseError
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "invalid request", response.Message)

}
