package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/venomuz/project2/storage/postgres"
	"github.com/venomuz/project2/storage/repo"
)

//IStorage ...
type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}
