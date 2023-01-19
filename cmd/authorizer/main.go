package main

import (
	"authorizer/api"
	pb "authorizer/proto"
	"github.com/praissik/web-app-engine/pkg/engine"
)

func init() {
	engine.InitEnv()
	engine.InitViper()
}

func main() {
	s, lis := engine.PrepareGrpcServer()

	pb.RegisterAuthServer(s, &api.Server{})

	engine.RunGrpcServer(s, lis)
}
