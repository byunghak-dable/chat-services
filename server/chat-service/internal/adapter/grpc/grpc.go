package grpc

import (
	"net"

	"github.com/widcraft/chat-service/internal/adapter/grpc/chat"
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/logger"
	"google.golang.org/grpc"
)

type Grpc struct {
	logger  logger.Logger
	server  *grpc.Server
	chatApp port.ChatApp
}

func New(logger logger.Logger, chatApp port.ChatApp) *Grpc {
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
