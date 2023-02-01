package router

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/di"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *Server) CreateChat(ctx context.Context, req *pb.Chat) (*pb.ChatResponse, error) {

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

	input := &entity.Chat{
		Content: req.Content,
	}

	// input := entity.Chat{
	// 	Id:        uint(req.Id),
	// 	Email:     req.Email,
	// 	CreatedAt: req.CreatedAt.AsTime(),
	// }

	tx, _ := db.Begin()
	h := di.InitializeChatHandler(tx, firebase)
	presenter, err := h.Create(input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println(h)
	fmt.Println(presenter)

	return &pb.ChatResponse{}, nil
}

func (s *Server) GetChat(ctx context.Context, req *pb.ChatRequest) (*pb.ChatResponse, error) {
	fmt.Println("Get")

	var (
		db = database.NewDB(config.Db{
			Host: config.DbHost,
		}, true)
		firebase = driver.NewFirebaseImpl(config.Firebase{})
	)

	tx, _ := db.Begin()
	h := di.InitializeChatHandler(tx, firebase)
	presenter, err := h.GetById(uint(req.Chat.Id))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println(h)
	fmt.Println(presenter)

	return &pb.ChatResponse{}, nil
}

func (s *Server) GetChatStream(req *pb.GetStreamRequest, server pb.ChatService_GetChatStreamServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := make(chan entity.Chat)

	go func() error {
		defer close(stream)
		var (
			db = database.NewDB(config.Db{
				Host: config.DbHost,
			}, true)
			firebase = driver.NewFirebaseImpl(config.Firebase{})
		)

		eg, _ := errgroup.WithContext(ctx)
		eg.Go(func() error {
			h := di.InitializeChatHandler(db, firebase)
			presenter, err := h.GetStream(uint(req.GroupId))
			if err != nil {
				return err
			}
			fmt.Println(presenter)
			return nil
		})
		if err := eg.Wait(); err != nil {
			return fmt.Errorf("failed to GetChatStreamService.Handle: %s", err)
		}

		return nil
	}()

	for {
		v := <-stream
		createdAt := timestamppb.New(v.CreatedAt)
		if err := server.Send(&pb.ChatResponse{
			Chat: &pb.Chat{
				From:      v.From,
				Content:   v.Content,
				CreatedAt: createdAt,
			},
		}); err != nil {
			return err
		}
	}
}
