package model

import "image"

type Movie struct {
	Name      string
	Mark      float64 `json:"mark"`
	Statistic int     `json:"statistic"`
	WantTOSee bool    `json:"want_to_see"`
	SeenOrNot bool    `json:"seen_or_not"`
	Photo     image.Image
	Brief     string `json:"brief"`
	Actors    string `json:"actors"`
	class     string `json:"class"`
	ShortComment
	FilmCritics
}
