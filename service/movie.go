package service

import (
	"movie/dao"
	"movie/model"
)

func SelectMovieByName(movieName string) (model.Movie, error) {
	return dao.SelectMovieByName(movieName)
}

func SelectMovieByNameDetail(movieName string) (model.Movie, error) {
	return dao.SelectMovieByNameDetail(movieName)
}

func AddShortComment(shortComment model.ShortComment) error {
	err := dao.AddShortComment(shortComment)
	return err
}

func AddUseful(MovieShortComment model.Movie) error {
	return dao.AddUseful(MovieShortComment)
}

func AddFilmCritics(critics model.FilmCritics) error {
	return dao.AddFilmCritics(critics)
}

func List() ([]string, error) {
	return dao.List()
}

func FindBYClass(class string) (model.Movie, error) {
	return dao.FindBYClass(class)
}

func SelectSelfIntroduce(username string) (introduce string, err error) {
	return dao.SelectSelfIntroduce(username)
}
