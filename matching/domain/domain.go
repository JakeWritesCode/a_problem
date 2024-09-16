package domain

import "a_problem/matching/store"

type DomainInterface interface {
	Seed()
}

type Domain struct {
	store store.PostgresStoreInterface
}

func NewDomain() *Domain {
	pgStore, err := store.NewPostgresStore()
	if err != nil {
		panic(err)
	}
	return &Domain{store: pgStore}
}
