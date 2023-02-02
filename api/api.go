package api

import (
	"authorizer/pkg/auth"
	"authorizer/pkg/proto/pb"
	"context"
	"log"
)

type Server struct {
	pb.UnimplementedAuthServer
}

func (s *Server) Register(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
	log.Printf("Received: %v: %v", in.GetEmail(), in.GetPassword())

	token, err := auth.Register(in.GetEmail(), in.GetPassword())
	if err != nil {
		return &pb.AuthReply{Status: 400, Message: err.Error()}, nil
	}
	return &pb.AuthReply{Status: 200, Message: token}, nil
}
