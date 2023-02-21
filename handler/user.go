package handler

import (
	"context"
	"fmt"

	// "github.com/hidenari-yuda/go-grpc-clean/domain/requests"
	// "github.com/hidenari-yuda/go-grpc-clean/domain/responses"
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

// server is used to implement helloworld.GreeterServer.
type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	UserInteractor interactor.UserInteractor
	Db             *database.Db
	Firebase       usecase.Firebase
}

func NewUserSercviceServer(userInteractor interactor.UserInteractor) *UserServiceServer {
	return &UserServiceServer{
		UserInteractor: userInteractor,
		Db:             database.NewDb(),
		Firebase:       driver.NewFirebaseImpl(),
	}
}

func (s *UserServiceServer) Create(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println("db is:", s.Db)
	fmt.Println("firebase is :", s.Firebase)

	// input, err := requests.NewUser(req)
	// if err != nil {
	// 	return nil, handleError(err)
	// }

	tx, _ := s.Db.Begin()
	res, err := s.UserInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return res, nil

}

func (s *UserServiceServer) GetById(ctx context.Context, req *pb.UserIdRequest) (*pb.User, error) {

	res, err := s.UserInteractor.GetById(uint(req.Id))
	if err != nil {
		return nil, handleError(err)
	}

	return res, nil
}
