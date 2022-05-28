package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/project2/genproto"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	return nil, nil
}
