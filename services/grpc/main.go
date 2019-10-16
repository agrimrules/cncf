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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
}

// Post is a basic struct indicating a user and message
type Post struct {
	gorm.Model
	User    string `json:"user"`
	Message string `json:"message"`
}

func dataToPB(data *Post) *message.Message {
	return &message.Message{
		User:    data.User,
		Message: data.Message,
	}
}

func (*server) CreateMessage(ctx context.Context, req *message.CreateMessageReq) (*message.CreateMessageRes, error) {
	msg := req.GetMessage()
	data := Post{
		User:    msg.GetUser(),
		Message: msg.GetMessage(),
	}
	db.Create(data)
	return &message.CreateMessageRes{Message: msg}, nil
}

func (*server) GetMessage(req *message.GetMessageReq, stream message.MessageService_GetMessageServer) error {
	user := req.GetUser()
	posts := []Post{}
	if err := db.Where("user LIKE ?", user).Find(&posts).Error; err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not fetch data for user %s error: %v", user, err),
		)
	}
	for _, p := range posts {
		stream.Send(&message.GetMessageRes{Message: dataToPB(&p)})
	}
	return nil
}

func (*server) ListMessage(req *message.ListMessageRequest, stream message.MessageService_ListMessageServer) error {
	posts := []Post{}
	if err := db.Find(&posts).Error; err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error retrieving data: %v", err),
		)
	}
	for _, p := range posts {
		stream.Send(&message.ListMessageResponse{Message: dataToPB(&p)})
	}
	return nil
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
		log.Fatal(err)
		panic("Failed to connect to DB")
	}

	// Default settings for GRPC is tcp on 50051
	listener, err := net.Listen(conf.PROTOCOL, ":"+conf.PORT)

	if err != nil {
		log.Fatalf("Unable to listen: %v", err)
	}

	options := []grpc.ServerOption{}

	s := grpc.NewServer(options...)

	// Bind service to Server
	message.RegisterMessageServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
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
