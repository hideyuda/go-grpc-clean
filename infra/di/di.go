package di

import (
	"context"

	"github.com/hidenari-yuda/go-grpc-clean/handler"
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/repository"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"
	"google.golang.org/grpc"
)
// RegisterServiceServer is a function to register service server
func RegisterServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	regsiterUserServiceServer(ctx, s, db, firebase)
	registerChatServiceServer(ctx, s, db, firebase)
}

// regsiterUserServiceServer is a function to register user service server
func regsiterUserServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	userRepository := repository.NewUserRepositoryImpl(db)
	pb.RegisterUserServiceServer(s, handler.NewUserSercviceServer(interactor.NewUserInteractorImpl(firebase, userRepository)))
}

// registerChatServiceServer is a function to register chat service server
func registerChatServiceServer(ctx context.Context, s *grpc.Server, db *database.Db, firebase usecase.Firebase) {
	chatRepository := repository.NewChatRepositoryImpl(db)
	// chatGroupRepository := repository.NewChatGroupRepositoryImpl(db)
	// chatUserRepository := repository.NewChatUserRepositoryImpl(db)
	pb.RegisterChatServiceServer(s, handler.NewChatSercviceServer(interactor.NewChatInteractorImpl(firebase, chatRepository)))
}
