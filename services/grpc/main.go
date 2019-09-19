package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/agrimrules/cncf/services/config"
	message "github.com/agrimrules/cncf/services/grpc/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
)

type MessageServiceServer struct {
}

type Post struct {
	gorm.Model
	User    string `json:"user"`
	Message string `json:"message"`
}

func (s *MessageServiceServer) CreateMessage(ctx context.Context, req *message.CreateMessageReq) (*message.CreateMessageRes, error) {
	msg := req.GetMessage()
	data := Post{
		User:    msg.GetUser(),
		Message: msg.GetMessage(),
	}
	db.Create(data)
	return &message.CreateMessageRes{Message: msg}, nil
}

func (s *MessageServiceServer) GetMessage(ctx context.Context, req *message.GetMessageReq) (*message.GetMessageRes, error) {
	// TODO: Actually implement this service
	return nil, nil
}

var db *gorm.DB

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		panic("Failed to load config")
	}
	db, err := gorm.Open(conf.DIALECT, conf.URL)
	db.AutoMigrate(&Post{})

	if err != nil {
		panic("Failed to connect to DB")
		log.Fatal(err)
	}

	// Default settings for GRPC is tcp on 50051
	listener, err := net.Listen(conf.PROTOCOL, ":"+conf.PORT)

	if err != nil {
		log.Fatal("Unable to listen: %v", err)
	}

	options := []grpc.ServerOption{}

	s := grpc.NewServer(options...)
	service := &MessageServiceServer{}

	// Bind service to Server
	message.RegisterMessageServiceServer(s, service)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatal("Failed to serve: %v", err)
		}
	}()

	fmt.Println("Up and running")

	// Channel to recieve OS signals
	c := make(chan os.Signal)

	// Listening for Interrupts
	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("recieved interrupt")
	s.Stop()
	defer db.Close()

}
