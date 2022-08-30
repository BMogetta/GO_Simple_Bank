package gapi

import (
	db "github.com/BMogetta/Simple_bank/postgres/sqlc"
	pb "github.com/BMogetta/Simple_bank/proto_go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// convertUser takes a db User object and returns a pb User, filtering sensitive information
func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
