package handler

import (
	"context"

	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"
)

type ChatGroupServiceServer struct {
	pb.UnimplementedChatGroupServiceServer
	ChatGroupInteractor interactor.ChatGroupInteractor
	Db                  *database.Db
	Firebase            usecase.Firebase
}

func NewChatGroupSercviceServer(chatGroupInteractor interactor.ChatGroupInteractor) *ChatGroupServiceServer {
	return &ChatGroupServiceServer{
		ChatGroupInteractor: chatGroupInteractor,
		Db:                  database.NewDb(),
		Firebase:            driver.NewFirebaseImpl(),
	}
}

// create chat group
func (s *ChatGroupServiceServer) Create(ctx context.Context, req *pb.ChatGroup) (*pb.ChatGroup, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ChatGroupInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return res, nil
}

// update chat group
func (s *ChatGroupServiceServer) Update(ctx context.Context, req *pb.ChatGroup) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ChatGroupInteractor.Update(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// delete chat group
func (s *ChatGroupServiceServer) Delete(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ChatGroupInteractor.Delete(uint(req.Id))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// get chat group by id
func (s *ChatGroupServiceServer) GetById(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatGroup, error) {

	res, err := s.ChatGroupInteractor.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get chat group by user id
func (s *ChatGroupServiceServer) GetListByUserId(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatGroupList, error) {

	res, err := s.ChatGroupInteractor.GetListByUserId(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ChatGroupList{ChatGroupList: res}, nil
}
