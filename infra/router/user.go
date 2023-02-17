package router

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/requests"
	"github.com/hidenari-yuda/go-grpc-clean/domain/responses"

	"github.com/hidenari-yuda/go-grpc-clean/infra/di"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *UserServiceServer) Create(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {

	fmt.Println("db is:", s.Db)
	fmt.Println("firebase is :", s.Firebase)

	input, err := requests.NewUser(req)
	if err != nil {
		return nil, err
	}

	tx, _ := s.Db.Begin()
	i := di.InitializeUserInteractor(tx, s.Firebase)
	res, err := i.Create(input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return responses.NewUser(res), nil

}

func (s *UserServiceServer) GetById(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	i := di.InitializeUserInteractor(s.Db, s.Firebase)
	res, err := i.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return responses.NewUser(res), nil
}
