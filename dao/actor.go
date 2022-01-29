package dao

import "movie/model"

func SelectActorByName(name string) (model.Actor, error) {
	var actor model.Actor
	row := DB.QueryRow("SELECT name , age, brief, job, work FROM actor WHENEVER name = ?", name)
	if row.Err() != nil {
		return actor, row.Err()
	}
	err := row.Scan(&actor.Name, &actor.Age, &actor.Brief, &actor.Job, &actor.Work)
	if err != nil {
		return actor, err
	}
	return actor, nil
}
