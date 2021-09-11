package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "github.com/nrnrk/bazel-sample/hello/hello"
)

const (
	address        = "localhost:55550"
	requestMessage = "Hey"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		hoge()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func hoge() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := client.Say(ctx, &pb.HelloRequest{Message: requestMessage})
	if err != nil {
		log.Fatalf("failed to say hello: %v", err)
	}
	log.Printf("Hello message: %s", r.GetMessage())
}
