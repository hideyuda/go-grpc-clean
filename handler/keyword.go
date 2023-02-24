package handler

import (
	"context"

	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"
)

type KeywordServiceServer struct {
	pb.UnimplementedKeywordServiceServer
	KeywordInteractor interactor.KeywordInteractor
	Db                *database.Db
	Firebase          usecase.Firebase
}

func NewKeywordSercviceServer(KeywordInteractor interactor.KeywordInteractor) *KeywordServiceServer {
	return &KeywordServiceServer{
		KeywordInteractor: KeywordInteractor,
		Db:                database.NewDb(),
		Firebase:          driver.NewFirebaseImpl(),
	}
}

// create chat group
func (s *KeywordServiceServer) Create(ctx context.Context, req *pb.Keyword) (*pb.Keyword, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.KeywordInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return res, nil
}

// update chat group
func (s *KeywordServiceServer) Update(ctx context.Context, req *pb.Keyword) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.KeywordInteractor.Update(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// delete chat group
func (s *KeywordServiceServer) Delete(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.KeywordInteractor.Delete(req.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// get chat group by id
func (s *KeywordServiceServer) GetById(ctx context.Context, req *pb.ChatIdRequest) (*pb.Keyword, error) {

	res, err := s.KeywordInteractor.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get chat group by user id
func (s *KeywordServiceServer) GetListByUserId(ctx context.Context, req *pb.ChatIdRequest) (*pb.KeywordList, error) {

	res, err := s.KeywordInteractor.GetListByUserId(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.KeywordList{KeywordList: res}, nil
}
