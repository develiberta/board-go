package model

import "time"

type Post struct {
	Seq          int       `json:"seq,omitempty"`
	Id           string    `json:"id,omitempty"`
	Title        string    `json:"title"`
	Contents     string    `json:"contents"`
	Creator      string    `json:"creator"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedTime time.Time `json:"modified_time"`
}

var Posts map[int]Post
var LastSeq int
