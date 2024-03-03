package grpc

import (
	"messenger-service/internal/adapter/driving/grpc/chat"
	"messenger-service/internal/adapter/driving/grpc/chat/pb"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
	"net"

	"google.golang.org/grpc"
)

type Grpc struct {
	logger       driven.LoggerPort
	server       *grpc.Server
	messengerApp driving.MessengerServicePort
	port         string
}

func New(logger driven.LoggerPort, messenger driving.MessengerServicePort, port string) *Grpc {
	server := grpc.NewServer()
	pb.RegisterChatServer(server, chat.New(logger, messenger))

	return &Grpc{
		logger:       logger,
		messengerApp: messenger,
		server:       server,
		port:         port,
	}
}

func (g *Grpc) Run() error {
	listener, err := net.Listen("tcp", ":"+g.port)

	if err != nil {
		return err
	}

	return g.server.Serve(listener)
}

func (g *Grpc) Close() error {
	g.server.GracefulStop()
	return nil
}
