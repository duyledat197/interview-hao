
package deliveries

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/internal/services"
	"github.com/duyledat197/go-gen-tools/pb"
	"github.com/duyledat197/go-gen-tools/transform"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userDelivery struct {
	userService services.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserDelivery(userService services.UserService) pb.UserServiceServer {
	return &userDelivery{
		userService: userService,
	}
}

func (d *userDelivery) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	if err := d.userService.Create(ctx, transform.PbToUserPtr(req.GetUser())); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.CreateUserResponse{}, nil
}
