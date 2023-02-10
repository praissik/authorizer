package main

import (
	"authorizer/api"
	"authorizer/pkg/proto/pb"
	"github.com/praissik/web-app-engine/engine"
)

func init() {
	engine.InitEnv()
	engine.InitViper()
}

func main() {
	s, lis := engine.PrepareGrpcServer()

	pb.RegisterAuthServer(s, api.NewServer())

	engine.RunGrpcServer(s, lis)
}
