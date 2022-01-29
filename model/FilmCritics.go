package model

import "time"

type FilmCritics struct {
	movieName   string    `json:"movie_name"`
	Username    string    `json:"username"`
	Critics     string    `json:"critics"`
	CriticsTime time.Time `json:"critics_time"`
	Useful      int       `json:"useful"`
	Useless     int       `json:"useless"`
	Node
}
