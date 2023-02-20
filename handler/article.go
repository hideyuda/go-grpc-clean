package handler

import (
	"context"

	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ArticleServiceServer struct {
	pb.UnimplementedArticleServiceServer
	// ArticleInteractor interactor.ArticleInteractor
	Db       *database.Db
	Firebase usecase.Firebase
}

func (s *ArticleServiceServer) Create(ctx context.Context, req *pb.Article) (*pb.ArticleResponse, error) {

	// tx, _ := s.Db.Begin()
	// res, err := s.ArticleInteractor.Create(req)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, handleError(err)
	// }

	// tx.Commit()

	// return responses.NewArticle(res), nil
	return nil, nil
}

func (s *ArticleServiceServer) CreateArticleByKeyword(ctx context.Context, req *pb.ArticleRequest) (*pb.ArticleResponse, error) {

	// tx, _ := s.Db.Begin()
	// res, err := s.ArticleInteractor.Create(req)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, handleError(err)
	// }

	// tx.Commit()

	// return responses.NewArticle(res), nil
	return nil, nil
}
