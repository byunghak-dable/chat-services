package grpc

import (
	"messenger-service/internal/adapter/driven/config"
	"messenger-service/internal/adapter/driving/grpc/messenger"
	"messenger-service/internal/adapter/driving/grpc/messenger/pb"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
	"net"

	"google.golang.org/grpc"
)

type Grpc struct {
	logger driven.Logger
	server *grpc.Server
	port   string
}

func New(configStore *config.Store, logger driven.Logger, messengerApp driving.Messenger) *Grpc {
	server := grpc.NewServer()
	pb.RegisterChatServer(server, messenger.New(logger, messengerApp))

	return &Grpc{
		logger: logger,
		server: server,
		port:   configStore.GetGrpcPort(),
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
