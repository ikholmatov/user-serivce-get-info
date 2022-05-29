package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	pb "github.com/venomuz/project2/genproto"
	l "github.com/venomuz/project2/pkg/logger"
	"github.com/venomuz/project2/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	id1 := uuid.NewV4()
	id2 := uuid.NewV4()
	req.ID = id1.String()
	req.Addr.ID = id2.String()
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Error while inserting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}
	return user, err
}
func (s *UserService) GetByID(ctx context.Context, id *pb.ID) (*pb.User, error) {
	user, err := s.storage.User().GetByID(id.UserID)
	if err != nil {
		s.logger.Error("Error while getting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}
	return user, err
}
