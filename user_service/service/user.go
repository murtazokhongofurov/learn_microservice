package service
import (
	"context"

	pb "github.com/three_serviceSimple/user_service/genproto/user"
	l "github.com/three_serviceSimple/user_service/pkg/logger"
	"github.com/three_serviceSimple/user_service/storage"

	grpcclient "github.com/three_serviceSimple/user_service/service/grpc_client"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcclient.GrpcClientI
}

func NewUserService(db *sqlx.DB, log l.Logger, client grpcclient.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.logger.Error("error insert users", l.Any("Error insert users", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return user, nil
}
func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUser) (*pb.User, error) {
	user, err := s.storage.User().GetUserById(req)
	if err != nil {
		s.logger.Error("error insert users", l.Any("Error insert users", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return user, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().UpdateUser(req)
	if err != nil {
		s.logger.Error("error update", l.Any("Error update users", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrogn, please check user info")
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.GetUser) (*pb.Empty, error) {
	user, err := s.storage.User().DeleteUser(req)
	if err != nil {
		s.logger.Error("error delete", l.Any("Error delete users", err))
		return &pb.Empty{}, status.Error(codes.Internal, "something went wrogn, please check user info")
	}
	return user, nil
}
