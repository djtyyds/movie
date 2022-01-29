package service

import (
	"movie/dao"
	"movie/model"
)

func SelectActorByName(name string) (model.Actor, error) {
	return dao.SelectActorByName(name)
}
