package repository

import "github.com/dyammarcano/fullcycle_clean_architecture/internal/entity"

func Must(repo entity.OrderRepository, err error) entity.OrderRepository {
	if err != nil {
		panic(err)
	}
	return repo
}
