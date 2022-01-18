package model

import "time"

type ShortComment struct {
	Username    string    `json:"username"`
	Txt         string    `json:"txt"`
	commentTime time.Time `json:"comment_time"`
	Useful      int       `json:"useful"`
}
