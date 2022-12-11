package grpc

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat"
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/port"
	"google.golang.org/grpc"
)

type Grpc struct {
	logger  *log.Logger
	server  *grpc.Server
	chatApp port.ChatApp
}

func New(logger *log.Logger, chatApp port.ChatApp) *Grpc {
	server := grpc.NewServer()
	pb.RegisterChatServer(server, chat.New(logger))

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
	defer func() {
		if err = listener.Close(); err != nil {
			g.logger.Errorf("failed to close tcp connection", err)
		}
	}()

	if err = g.server.Serve(listener); err != nil {
		g.logger.Errorf("serve grpc error: %s", err)
	}
}

func (g *Grpc) Close() {
	g.server.GracefulStop()
	g.logger.Infoln("shutting down grpc server")
}
