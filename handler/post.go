package handler

import (
	"awesomeProject/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ListPostHandler(w http.ResponseWriter, r *http.Request) {
	list := make([]model.Post, 0)
	for _, post := range model.Posts {
		list = append(list, post)
	}

	w.WriteHeader(http.StatusOK)

	setHeader(w)
	json.NewEncoder(w).Encode(list)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seq, _ := strconv.Atoi(vars["seq"])
	post, ok := model.Posts[seq]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		setHeader(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	setHeader(w)
	json.NewEncoder(w).Encode(post)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		setHeader(w)
		return
	}
	model.LastSeq++
	post.Seq = model.LastSeq
	model.Posts[model.LastSeq] = post
	w.WriteHeader(http.StatusCreated)
	setHeader(w)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		setHeader(w)
		return
	}

	vars := mux.Vars(r)
	seq, _ := strconv.Atoi(vars["seq"])
	model.Posts[seq] = post
	w.WriteHeader(http.StatusOK)
	setHeader(w)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seq, _ := strconv.Atoi(vars["seq"])
	_, ok := model.Posts[seq]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		setHeader(w)
		return
	}
	delete(model.Posts, seq)
	w.WriteHeader(http.StatusOK)
	setHeader(w)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
}
