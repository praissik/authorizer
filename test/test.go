package test

import (
	"authorizer/pkg/database"
	"authorizer/pkg/proto/pb"
	"context"
	"github.com/praissik/web-app-engine/engine"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

type Engine struct {
	Client    pb.AuthClient
	Email     string
	BusyEmail string
	Password  string
	Ctx       context.Context
}

func initEnv(t *testing.T) {
	t.Setenv("LAUNCH_MODE", "test")
}
func initViper() {
	engine.InitViper()
}

func GetEngine(t *testing.T, authServer pb.AuthServer) Engine {
	initEnv(t)
	initViper()

	ctx := context.Background()
	client, _ := server(ctx, authServer)

	return Engine{
		Client:    client,
		Email:     "new@email.com",
		BusyEmail: "busy@email.com",
		Password:  "password",
		Ctx:       ctx,
	}
}

func server(ctx context.Context, authServer pb.AuthServer) (pb.AuthClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterAuthServer(baseServer, authServer)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewAuthClient(conn)

	return client, closer
}

func (e Engine) CreateAccount() {
	mongoClient := database.GetMongoClient()
	_, err := mongoClient.
		Database("test").
		Collection("account").
		InsertOne(context.TODO(), bson.D{{"email", e.BusyEmail}})
	if err != nil {
		log.Fatal(err)
	}
}
