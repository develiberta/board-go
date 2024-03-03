package handler

import (
	"awesomeProject/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/posts", nil)
	handler := MakeWebHandler()
	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []model.Post
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Title)
	assert.Equal("bbb", list[1].Title)
}

func TestGetJsonHandler(t *testing.T) {
	assert := assert.New(t)

	var post model.Post
	handler := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/posts/1", nil)

	handler.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&post)
	assert.Nil(err)
	assert.Equal("aaa", post.Title)
}

func TestCreateJsonHandler(t *testing.T) {
	assert := assert.New(t)

	var post model.Post
	handler := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/posts",
		strings.NewReader(`{"Seq":3, "Title":"ccc", "Contents":"ccc", "Creator":"ccc"}`))

	handler.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/posts/3", nil)
	handler.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&post)
	assert.Nil(err)
	assert.Equal("ccc", post.Title)
}

func TestUpdateJsonHandler(t *testing.T) {
	assert := assert.New(t)

	var post model.Post
	handler := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/posts",
		strings.NewReader(`{"Seq":3, "Title":"ccc", "Contents":"ccc", "Creator":"ccc"}`))

	handler.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/posts/3", nil)
	handler.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&post)
	assert.Nil(err)
	assert.Equal("ccc", post.Title)
}

func TestDeleteJsonHandler(t *testing.T) {
	assert := assert.New(t)

	handler := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/posts/1", nil)

	handler.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/posts/1", nil)
	handler.ServeHTTP(res, req)

	assert.Equal(http.StatusNotFound, res.Code)
}
