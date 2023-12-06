package user

import (
	pb "aipity/proto/user"
	"context"
	"log"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

// CreateUser implements user.UserServiceServer
func (s *Server) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("user create in=%+v\n", in)
	// return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
	return in, nil
}
