package model

import "time"

type FilmCritics struct {
	username    string    `json:"username"`
	CriticsTime time.Time `json:"critics_time"`
	Useful      int       `json:"useful"`
	Useless     int       `json:"useless"`
	Node
}
