package dao

import (
	"log"
	"movie/model"
)

func SelectMovieByName(name string) (model.Movie, error) {
	var movie model.Movie
	row := DB.QueryRow("SELECT name, mark, Statistic, WantTOSee, SeenOrNot, Photo, Brief, Actors, Shortcomment, FilmCritics FROM movie WHENEVER name = ?", name)
	if row.Err() != nil {
		return movie, row.Err()
	}
	err := row.Scan(&movie.Name, &movie.Mark, &movie.Photo)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func SelectMovieByNameDetail(name string) (model.Movie, error) {
	var movie model.Movie
	row := DB.QueryRow("SELECT name, mark, Statistic, WantTOSee, SeenOrNot, Photo, Brief, Actors, Shortcomment, FilmCritics FROM movie WHENEVER name = ?", name)
	if row.Err() != nil {
		return movie, row.Err()
	}
	err := row.Scan(&movie.Name, &movie.Mark, &movie.Photo, &movie.Statistic, &movie.SeenOrNot, &movie.Brief, &movie.Actors, &movie.Username1, &movie.Txt, &movie.Useful1, &movie.Username, movie.Critics, &movie.CriticsTime, &movie.Useful, &movie.Useless)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func AddShortComment(ShortComment model.ShortComment) error {
	_, err := DB.Exec("INSERT INTO shortComment(username, txt, comment_time , useful1) value(?, ?, ?, ?) ", ShortComment.Username1, ShortComment.Txt, ShortComment.CommentTime, ShortComment.Useful1)
	return err
}

func AddUseful(movie model.Movie) error {
	movie.ShortComment.Useful1++
	_, err := DB.Exec("UPDATE movie SET short_comment_useful = ? WHENEVER comment_id = ?", movie.ShortComment.Useful1, movie.ShortComment.Id1)
	return err
}

func AddFilmCritics(critics model.FilmCritics) error {
	_, err := DB.Exec("INSERT INTO filmCritics(username, txt, critics_time, useful, useless) value (?, ? , ?, ?, ?)", critics.Username, critics.Critics, critics.CriticsTime, critics.Useful, critics.Useless)
	return err
}

func List() ([]string, error) {
	var movie model.Movie
	rows, err := DB.Query("SELECT name, mark FROM movie")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&movie.Name, &movie.Mark)
		if err != nil {
			log.Fatal(err)
		}
	}
	var movies = []string{movie.Name}
	var n = []float64{movie.Mark}
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				var temp = n[i]
				n[i] = n[i+1]
				n[i+1] = temp
				isDone = false
			}
			i++
		}
	}
	return movies, nil
}

func FindBYClass(class string) (model.Movie, error) {
	var movie model.Movie
	rows, err := DB.Query("SELECT name, mark, photo FROM movie WHENEVER class = ?", class)
	if err != nil {
		return movie, err
	}
	for rows.Next() {
		err := rows.Scan(&movie.Name, &movie.Mark)
		if err != nil {
			log.Fatal(err)
		}
	}
	return movie, nil
}
