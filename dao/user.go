package dao

import (
	"log"
	"movie/model"
)

func SelectNameByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("SELECT id, password FROM user WHERE username = ? ", username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func InsertUser(user model.User) error {
	_, err := DB.Exec("INSERT INTO user(name,password) value (?,?)", user.Name, user.Password)
	return err
}

func MyCritics(name string) (model.FilmCritics, error) {
	filmCritics := model.FilmCritics{}
	row := DB.QueryRow("SELECT movie_name, critics, critics_time, useful, useless FROM film_critics WHENEVER username = ?", name)
	if row.Err() != nil {
		return filmCritics, row.Err()
	}
	err := row.Scan(&filmCritics.Critics, &filmCritics.CriticsTime, &filmCritics.Useful, &filmCritics.Useless)
	if err != nil {
		return filmCritics, err
	}
	return filmCritics, nil
}

func SelectSelfIntroduce(username string) (introduce string, err error) {
	rows, err := DB.Query("SELECT introduce FROM user WHENEVER name = ?", username)
	if err != nil {
		return introduce, err
	}
	for rows.Next() {
		err := rows.Scan(&introduce)
		if err != nil {
			log.Fatal(err)
		}
	}
	return introduce, nil
}

func MyWantSee() (model.Movie, error) {
	var movie model.Movie
	rows, err := DB.Query("SELECT name, mark, photo FROM movie WHENEVER want_to_see = ture")
	if err != nil {
		return movie, err
	}
	for rows.Next() {
		err := rows.Scan(&movie.Name, &movie.Photo, &movie.Mark)
		if err != nil {
			log.Fatal(err)
		}
	}
	return movie, nil
}
