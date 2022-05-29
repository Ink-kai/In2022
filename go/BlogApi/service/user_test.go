package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	router := Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Equal(t, "welcome to my site", response["message"])
}
func TestAddUser(t *testing.T) {
	test := []struct {
		name string
		data map[string]string
	}{
		{"测试1", map[string]string{"name": "刘凯", "password": "lje4wo5i86y7u845uy", "email": "123@21.chjk"}},
		{"测试2", map[string]string{"name": "ink", "password": "4444444", "email": "b.zweeucm@hirwixktku.aw"}},
		{"测试3", map[string]string{"name": "faew/", "password": "423423", "email": "123@qq.com"}},
	}
	r := Setup()
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			jsonByte, _ := json.Marshal(tt.data)
			req := httptest.NewRequest("POST", "/api/test/AddUser", bytes.NewReader(jsonByte))
			req.Header.Add("Cookie", "mode=test")
			req.Header.Add("Content-Type", "application/form-data; charset=utf-8")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
			var data map[string]map[string]interface{}
			json.Unmarshal([]byte(w.Body.String()), &data)
			assert.Equal(t, tt.data["name"], data["message"]["name"])
			assert.Equal(t, tt.data["email"], data["message"]["email"])
		})
	}
}
