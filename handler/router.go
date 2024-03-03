package handler

import (
	"awesomeProject/model"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"time"
)

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

func MakeWebHandler() http.Handler {
	router := mux.NewRouter()

	// 아래에 새로운 핸들러 등록
	router.HandleFunc("/posts", ListPostHandler).Methods(METHOD_GET)
	router.HandleFunc("/posts/{seq:[0-9]+}", GetPostHandler).Methods(METHOD_GET)
	router.HandleFunc("/posts", CreatePostHandler).Methods(METHOD_POST)
	router.HandleFunc("/posts/{seq:[0-9]+}", UpdatePostHandler).Methods(METHOD_PUT)
	router.HandleFunc("/posts/{seq:[0-9]+}", DeletePostHandler).Methods(METHOD_DELETE)

	model.Posts = make(map[int]model.Post)
	model.Posts[1] = model.Post{Seq: 1, Title: "title1", Contents: "contents1", Creator: "creator1", CreatedTime: time.Now(), ModifiedTime: time.Now()}
	model.Posts[2] = model.Post{Seq: 2, Title: "title2", Contents: "contents2", Creator: "creator2", CreatedTime: time.Now(), ModifiedTime: time.Now()}

	model.LastSeq = 2

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{METHOD_GET, METHOD_POST, METHOD_PUT, METHOD_DELETE},
	})

	return c.Handler(router)
}
