package handler

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *ChatServiceServer) Create(ctx context.Context, req *pb.Chat) (*pb.ChatResponse, error) {

	// Convert context.Context to echo.Context in gRPC server

	fmt.Println("Create")

	input := &entity.Chat{
		From:    uint(req.From),
		Content: req.Content,
	}

	tx, _ := s.Db.Begin()
	res, err := s.ChatInteractor.Create(input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatResponse{
		Chat: &pb.Chat{
			Id:        uint32(res.Id),
			From:      uint32(res.From),
			Content:   res.Content,
			CreatedAt: timestamppb.New(res.CreatedAt),
		},
	}, nil
}

func (s *ChatServiceServer) GetById(ctx context.Context, req *pb.GetChatByIdRequest) (*pb.ChatResponse, error) {
	fmt.Println("Get")

	// var (
	// 	db       = database.NewDb()
	// 	firebase = driver.NewFirebaseImpl()
	// )

	res, err := s.ChatInteractor.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ChatResponse{
		Chat: &pb.Chat{
			Id:        uint32(res.Id),
			From:      uint32(res.From),
			Content:   res.Content,
			CreatedAt: timestamppb.New(res.CreatedAt),
		},
	}, nil
}

func (s *ChatServiceServer) GetStream(req *pb.GetChatStreamRequest, server pb.ChatService_GetStreamServer) error {
	fmt.Println("GetStream")
	// h := di.InitializeChatHandler(s.Db, s.Firebase)
	// err := h.GetStream(req *pb.GetStreamRequest, server pb.ChatService_GetChatStreamServer)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := make(chan entity.Chat)

	go func() {
		err := s.ChatInteractor.GetStream(ctx, stream)
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		v := <-stream
		createdAt := timestamppb.New(v.CreatedAt)
		if err := server.Send(&pb.ChatResponse{
			Chat: &pb.Chat{
				From:      uint32(v.From),
				Content:   v.Content,
				CreatedAt: createdAt,
			},
		}); err != nil {
			return err
		}
	}
}
