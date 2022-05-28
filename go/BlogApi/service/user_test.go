package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/liuhongdi/gintest01/pkg/unittest"
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
		{"测试1", map[string]string{"name": "刘凯", "password": "lje4wo5i86y7u845uy"}},
		{"测试2", map[string]string{"name": "ink", "password": "", "email": "1@q.com"}},
	}
	r := Setup()
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			var w *httptest.ResponseRecorder
			w = unittest.PostForm("/api/test/AddUser", tt.data, r)
			assert.Equal(t, http.StatusOK, w.Code)
			body, _ := ioutil.ReadAll(w.Body)
			var data map[string]map[string]interface{}
			json.Unmarshal(body, &data)
			assert.Equal(t, tt.data["name"], data["message"]["name"])
		})
	}
}
