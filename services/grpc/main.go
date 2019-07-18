package main

import (
  "os"
  "os/signal"
  "context"
  "fmt"
  "log"
  "net"
  message "github.com/agrimrules/cncf/services/grpc/proto"
  "google.golang.org/grpc"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type MessageServiceServer struct {
}

type Post struct {
  gorm.Model
  User string `json:"user"`
  Message string  `json:"message"`
}

func (s *MessageServiceServer) CreateMessage(ctx context.Context, req *message.CreateMessageReq) (*message.CreateMessageRes, error) {
  msg := req.GetMessage();
  data := Post{
    User: msg.GetUser(),
    Message: msg.GetMessage(),
  }
  db.Create(data)
  return &message.CreateMessageRes{Message: msg}, nil
}

func (s *MessageServiceServer) GetMessage(ctx context.Context, req *message.GetMessageReq) (*message.GetMessageRes, error)  {
  // TODO: Actually implement this service
  return nil, nil
}

var db *gorm.DB

func main()  {
  db, err := gorm.Open("sqlite3", "data.db")
  db.AutoMigrate(&Post{})

  if err != nil {
    panic("Failed to connect to DB")
    log.Fatal(err)
  }

  // Default settings for GRPC is tcp on 50051
  listener, err := net.Listen("tcp", ":50051")

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
