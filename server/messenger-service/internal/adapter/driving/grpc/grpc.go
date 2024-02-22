package grpc

import (
	"net"

	"github.com/widcraft/messenger-service/internal/adapter/driving/grpc/chat"
	"github.com/widcraft/messenger-service/internal/adapter/driving/grpc/chat/pb"
	"github.com/widcraft/messenger-service/internal/port/driven"
	"github.com/widcraft/messenger-service/internal/port/driving"
	"google.golang.org/grpc"
)

type Grpc struct {
	logger       driven.LoggerPort
	server       *grpc.Server
	messengerApp driving.MessengerServicePort
}

func New(logger driven.LoggerPort, chatApp driving.MessengerServicePort) *Grpc {
	server := grpc.NewServer()
	pb.RegisterChatServer(server, chat.New(logger, chatApp))

	return &Grpc{
		logger:       logger,
		messengerApp: chatApp,
		server:       server,
	}
}

func (g *Grpc) Run(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		g.logger.Errorf("failed to listen on port %s", port)
	}

	if err = g.server.Serve(listener); err != nil {
		g.logger.Errorf("serve grpc error: %s", err)
	}
}

func (g *Grpc) OnExit() error {
	g.server.GracefulStop()
	return nil
}
