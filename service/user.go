package service

import (
	"database/sql"
	"movie/dao"
	"movie/model"
)

func IsRepeatName(username string) (bool, error) {
	_, err := dao.SelectNameByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsPasswordPlausible(user model.User) bool {
	if len(user.Password) < 8 {
		return false
	}
	return true
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}

func IsPasswordRight(username string, password string) (bool, error) {
	user, err := dao.SelectNameByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}
