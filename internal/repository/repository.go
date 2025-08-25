package repository

import "github.com/dyammarcano/fullcycle_clean_architecture/internal/domain"

func Must(repo domain.OrderRepository, err error) domain.OrderRepository {
	if err != nil {
		panic(err)
	}

	return repo
}
