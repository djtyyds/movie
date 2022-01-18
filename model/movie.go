package model

import "image"

type Movie struct {
	Mark      float64 `json:"mark"`
	Statistic int     `json:"statistic"`
	WantTOSee bool    `json:"want_to_see"`
	SeenOrNot bool    `json:"seen_or_not"`
	Photo     image.Image
	brief     string `json:"brief"`
	actors    string `json:"actors"`
	shortComment
	filmCritics
}
