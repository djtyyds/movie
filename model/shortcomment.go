package model

import "time"

type ShortComment struct {
	Id1         string    `json:"id"`
	Username1   string    `json:"username"`
	Txt         string    `json:"txt"`
	CommentTime time.Time `json:"comment_time"`
	Useful1     int       `json:"useful"`
}
