package handler

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/requests"
	"github.com/hidenari-yuda/go-grpc-clean/domain/responses"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *UserServiceServer) Create(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {

	fmt.Println("db is:", s.Db)
	fmt.Println("firebase is :", s.Firebase)

	input, err := requests.NewUser(req)
	if err != nil {
		return nil, handleError(err)
	}

	tx, _ := s.Db.Begin()
	res, err := s.UserInteractor.Create(input)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return responses.NewUser(res), nil

}

func (s *UserServiceServer) GetById(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	res, err := s.UserInteractor.GetById(uint(req.Id))
	if err != nil {
		return nil, handleError(err)
	}

	return responses.NewUser(res), nil
}
