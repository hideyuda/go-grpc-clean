package responses

import (
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

func NewUser(res *pb.User) *pb.UserResponse {
	return &pb.UserResponse{
		Error: false,
		User: res,
	}
}

// func NewUserList (res []*pb.User) *pb.UserListResponse {
// 	return &pb.UserListResponse{
// 		Error: false,
// 		UserList: res,
// 	}
// }

// func NewUserList(ress []*entity.User) []*pb.User {
// 	var resList []*pb.User
// 	for _, res := range ress {
// 		resList = append(resList, NewUser(res))
// 	}
// 	return resList
// }
