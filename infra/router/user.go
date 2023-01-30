package router

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"

	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/di"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {

	// Convert context.Context to echo.Context in gRPC server

	fmt.Println("Create")

	var (
		db = database.NewDB(config.Db{
			Host: config.DbHost,
		}, true)
		firebase = driver.NewFirebaseImpl(config.Firebase{})
	)

	// err := bindAndValidate(c, req)
	// if err != nil {
	// 	return nil, err
	// }

	input := &entity.SignUpParam{
		Email: req.Email,
	}

	// input := entity.User{
	// 	Id:        uint(req.Id),
	// 	Email:     req.Email,
	// 	CreatedAt: req.CreatedAt.AsTime(),
	// }

	tx, _ := db.Begin()
	h := di.InitializeUserHandler(tx, firebase)
	presenter, err := h.SignUp(input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println(h)
	fmt.Println(presenter)

	return &pb.UserResponse{}, nil
}

func (s *Server) GetUser(context.Context, *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("Create")

	var (
		db = database.NewDB(config.Db{
			Host: config.DbHost,
		}, true)
		firebase = driver.NewFirebaseImpl(config.Firebase{})
	)

	tx, _ := db.Begin()
	h := di.InitializeUserHandler(tx, firebase)
	// presenter, err := h.SignUp(req)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }
	tx.Commit()
	fmt.Println(h)
	// fmt.Println(presenter)

	return &pb.UserResponse{}, nil
}
