package grpc

import (
	"net"

	"github.com/widcraft/chat-service/internal/adapter/primary/grpc/chat"
	"github.com/widcraft/chat-service/internal/adapter/primary/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/infra"
	"github.com/widcraft/chat-service/internal/port"
	"google.golang.org/grpc"
)

type Grpc struct {
	logger  infra.Logger
	server  *grpc.Server
	chatApp port.MessageService
}

func New(logger infra.Logger, chatApp port.MessageService) *Grpc {
	server := grpc.NewServer()
	pb.RegisterChatServer(server, chat.New(logger, chatApp))

	return &Grpc{
		logger:  logger,
		chatApp: chatApp,
		server:  server,
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

func (g *Grpc) Close() error {
	g.server.GracefulStop()
	return nil
}
