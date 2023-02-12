package api

import (
	"authorizer/pkg/auth"
	"authorizer/pkg/proto/pb"
	"context"
	"log"
)

type server struct {
	pb.UnimplementedAuthServer
}

func NewServer() pb.AuthServer {
	return &server{}
}

func (s *server) Register(ctx context.Context, in *pb.AuthRequest) (*pb.AuthReply, error) {
	log.Printf("Received: %v", in.GetEmail())

	token, err := auth.Register(in.GetEmail(), in.GetPassword())
	if err != nil {
		return &pb.AuthReply{Status: 400, Message: err.Error()}, nil
	}
	return &pb.AuthReply{Status: 200, Message: token}, nil
}
