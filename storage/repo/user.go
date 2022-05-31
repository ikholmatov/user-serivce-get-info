package repo

import (
	pb "github.com/venomuz/project2/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	GetByID(ID string) (*pb.User, error)
	DeleteByID(ID string) (*pb.GetIdFromUser, error)
}
