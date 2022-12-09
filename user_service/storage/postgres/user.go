package postgres

import (
	pb "github.com/three_serviceSimple/user_service/genproto/user"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	userResp := pb.User{}
	id := uuid.New().String()
	err := r.db.QueryRow(`
	INSERT INTO users (
		id,
		full_name, 
		bio,
		phone_number,
		password
		) 
	VALUES ($1,$2,$3,$4,$5) 
		RETURNING 
		id, full_name,bio, phone_number, password`,
		id, user.FullName, user.Bio, user.PhoneNumber, user.Password,
	).Scan(
		&id,
		&userResp.FullName,
		&userResp.Bio,
		&userResp.PhoneNumber,
		&userResp.Password,
	)
	if err != nil {
		return &pb.User{}, err
	}
	return &userResp, nil
}
func (r *userRepo) GetUserById(req *pb.GetUser) (*pb.User, error) {
	response := pb.User{}
	err := r.db.QueryRow(`select 
			id, 
			full_name,
			bio,
			phone_number,
			password
			from users 
			where id = $1`, req.Id).Scan(
		&response.Id, &response.FullName,
		&response.Bio, &response.PhoneNumber,
		&response.Password)
	if err != nil {
		return &pb.User{}, err
	}
	return &response, nil
}
func (r *userRepo) UpdateUser(req *pb.User) (*pb.User, error) {
	user := pb.User{}
	err := r.db.QueryRow(`Update  users set 
			full_name = $1,
			bio = $2,
			phone_number = $3,
			password = $4
			where id = $5
			RETURNING
			id,
			full_name,
			bio,
			phone_number,
			password
			`, req.FullName,
		req.Bio,
		req.PhoneNumber,
		req.Password,
		req.Id).Scan(
		&user.Id,
		&user.FullName,
		&user.Bio,
		&user.PhoneNumber,
		&user.Password,
	)
	if err != nil {
		return &pb.User{}, err
	}
	return &user, nil
}
func (r *userRepo) DeleteUser(req *pb.GetUser) (*pb.Empty, error) {

	_, err := r.db.Exec(`delete from users where id = $1`, req.Id)
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil

}
