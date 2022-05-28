package repo

import (
	pb "github.com/venomuz/project2/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
}