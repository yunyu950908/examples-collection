package user

import (
	"context"
	userpb "grpc_example/proto/gen/go"
)

type Service struct {
	userpb.UnimplementedUserServiceServer
}

func (s *Service) GetUser(c context.Context, req *userpb.GetUserRequest) (res *userpb.GetUserResponse, err error) {
	return &userpb.GetUserResponse{
		Id: req.Id,
		User: &userpb.User{
			Name:       "Todd",
			Age:        18,
			IsVerified: true,
			Gender:     userpb.Gender_MALE,
		},
	}, err
}
