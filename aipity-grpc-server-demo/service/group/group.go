package group

import (
	pb "aipity/proto/group"
	"context"
	"log"
)

type Server struct {
	pb.UnimplementedGroupServiceServer
}

// CreateUser implements group.GroupServiceServer
func (s *Server) CreateGroup(ctx context.Context, in *pb.Group) (*pb.Group, error) {
	log.Printf("group create in=%+v\n", in)
	return in, nil
}
