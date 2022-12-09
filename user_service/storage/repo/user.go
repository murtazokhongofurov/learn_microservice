package repo

import (
	pb "github.com/three_serviceSimple/user_service/genproto/user"
)

type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
	GetUserById(*pb.GetUser) (*pb.User, error)
	UpdateUser(*pb.User) (*pb.User,error)
	DeleteUser(*pb.GetUser) (*pb.Empty, error)
}
