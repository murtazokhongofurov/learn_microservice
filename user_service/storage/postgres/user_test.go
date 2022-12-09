package postgres
import (
	"github.com/three_serviceSimple/user_service/config"
	pb "github.com/three_serviceSimple/user_service/genproto/user"
	"github.com/three_serviceSimple/user_service/pkg/db"
	"github.com/three_serviceSimple/user_service/storage/repo"

	"testing"

	"github.com/stretchr/testify/suite"
)

type UserSuiteTest struct {
	suite.Suite
	CleanUpfunc func()
	Repository  repo.UserStorageI
}

func (s *UserSuiteTest) SetubSuite() {
	pgPool, cleanUpfunc := db.ConnectToDBForSuite(config.Load())
	s.Repository = NewUserRepo(pgPool)
	s.CleanUpfunc = cleanUpfunc
}

func (s *UserSuiteTest) TestUserCrub() {
	userCreate := pb.User{
		FullName:    "suite full_name",
		Bio:         "suite bio",
		PhoneNumber: "suite phone_number",
		Password:    "suite password",
	}
	user, err := s.Repository.CreateUser(&userCreate)
	s.Nil(err)
	s.NotEmpty(user)
	updateUser := pb.User{
		Id:          user.Id,
		FullName:    "update new full_name",
		Bio:         "update bio",
		PhoneNumber: "update phone_number",
		Password:    "update email",
	}
	userUpdate, err := s.Repository.UpdateUser(&updateUser)
	s.Nil(err)
	s.NotEmpty(userUpdate)

	getUser, err := s.Repository.GetUserById(&pb.GetUser{Id: userUpdate.Id})
	s.Nil(err)
	s.NotEmpty(getUser)
	s.Equal(userUpdate.FullName, getUser.FullName)

	deleteUser, err := s.Repository.DeleteUser(&pb.GetUser{Id: getUser.Id})
	s.Nil(err)
	s.NotEmpty(deleteUser)
}

func (s *UserSuiteTest) TearDownSuite() {
	s.CleanUpfunc()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuiteTest))
}
