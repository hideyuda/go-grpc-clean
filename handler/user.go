package handler

import (
	"context"

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

// create user
func (s *UserServiceServer) Create(ctx context.Context, req *pb.User) (*pb.User, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, handleError(err)
	}

	res, err := s.UserInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return res, nil

}

// update user
// func (s *UserServiceServer) Update(ctx context.Context, req *pb.User) (*pb.User, error) {

// 	tx, err := s.Db.Begin()
// 	if err != nil {
// 		return nil, handleError(err)
// 	}

// 	res, err := s.UserInteractor.Update(req)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, handleError(err)
// 	}

// 	tx.Commit()

// 	return res, nil
// }

// delete user
// func (s *UserServiceServer) Delete(ctx context.Context, req *pb.UserIdRequest) (*pb.User, error) {

// 	tx, err := s.Db.Begin()
// 	if err != nil {
// 		return nil, handleError(err)
// 	}

// 	// res, err := s.UserInteractor.Delete(uint(req.Id))
// 	// if err != nil {
// 	// 	tx.Rollback()
// 	// 	return nil, handleError(err)
// 	// }

// 	tx.Commit()

// 	return res, nil
// }


func (s *UserServiceServer) GetById(ctx context.Context, req *pb.UserIdRequest) (*pb.User, error) {

	res, err := s.UserInteractor.GetById(uint(req.Id))
	if err != nil {
		return nil, handleError(err)
	}

	return res, nil
}
