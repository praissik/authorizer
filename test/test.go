package test

import (
	"authorizer/pkg/db"
	"authorizer/pkg/proto/pb"
	"context"
	"github.com/praissik/web-app-engine/engine"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

type Engine struct {
	Ctx    context.Context
	Client pb.AuthClient
}

func initEnv(t *testing.T) {
	t.Setenv("LAUNCH_MODE", "test")
}
func initViper() {
	engine.InitViper()
}

func cleanDatabase() error {
	mongoClient, deferF, err := db.GetMongoClient()
	defer deferF()
	if err != nil {
		return err
	}

	collections, _ := mongoClient.Database(viper.GetString("mongo.db")).ListCollectionNames(
		context.TODO(),
		bson.D{})
	if err != nil {
		return err
	}
	for _, collection := range collections {
		err = mongoClient.Database(viper.GetString("mongo.db")).Collection(collection).Drop(context.TODO())
		if err != nil {
			return err
		}
	}
	return nil
}

func GetEngine(t *testing.T, authServer pb.AuthServer) (*Engine, error) {
	initEnv(t)
	initViper()

	ctx := context.Background()
	client, _ := server(ctx, authServer)

	if err := cleanDatabase(); err != nil {
		return nil, err
	}

	return &Engine{
		Ctx:    ctx,
		Client: client,
	}, nil
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
