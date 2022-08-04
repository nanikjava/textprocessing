package repository

import "rockt/repository/model"

//Repository define the different operation that a repository should perform
type Repository interface {
	//Close closes the repository
	Close()

	//Create create all the necessary things related to the repository
	Create() error

	//BulkInsert performs bulk insertion to the repository
	BulkInsert(log []model.Datarecord) error

	Query(from string, to string) []model.Datarecord
}
