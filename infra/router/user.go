package router

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/di"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {

	// Convert context.Context to echo.Context in gRPC server

	fmt.Println("Create")

	var (
		db       = database.NewDb()
		firebase = driver.NewFirebaseImpl()
	)

	input := &entity.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: req.CreatedAt.AsTime(),
	}

	tx, _ := db.Begin()
	i := di.InitializeUserInteractor(tx, firebase)
	res, err := i.Create(input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.UserResponse{
		Error: false,
		User: &pb.User{
			Id:        uint32(res.Id),
			Name:      res.Name,
			Email:     res.Email,
			Password:  res.Password,
			CreatedAt: timestamppb.New(res.CreatedAt),
		},
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("Get")

	var (
		db       = database.NewDb()
		firebase = driver.NewFirebaseImpl()
	)

	tx, _ := db.Begin()
	i := di.InitializeUserInteractor(tx, firebase)
	res, err := i.GetById(uint(req.Id))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.UserResponse{
		Error: false,
		User: &pb.User{
			Id:        uint32(res.Id),
			Name:      res.Name,
			Email:     res.Email,
			Password:  res.Password,
			CreatedAt: timestamppb.New(res.CreatedAt),
		},
	}, nil
}
