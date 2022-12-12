package grpc

import (
	log "github.com/sirupsen/logrus"
)

type Grpc struct {
	logger  *log.Logger
	server  *grpc.Server
	userApp port.userApp
}

func New(logger *log.Logger, userApp port.userApp) *Grpc {
	server := grpc.NewServer()
	pb.RegisterChatServer(server, chat.New(logger))

	return &Grpc{
		logger:  logger,
		userApp: userApp,
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
