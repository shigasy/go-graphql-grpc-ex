package main

import (
	"log"
	"net"

	"github.com/shigasy/go-graphql-grpc-ex/article/pb"
	"github.com/shigasy/go-graphql-grpc-ex/article/repository"
	"github.com/shigasy/go-graphql-grpc-ex/article/service"
	"google.golang.org/grpc"
)

func main() {

	// articleサーバーに接続
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	// Repositoryを作成
	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	// Serviceを作成
	service := service.NewService(repository)

	//サーバーにarticleサービスを登録
	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	//articleサーバーを起動
	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
