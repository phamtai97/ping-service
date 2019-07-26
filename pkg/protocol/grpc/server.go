package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"
	pingservice_proto "pingservice/pkg/api"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//RunServer run gRPC service
func RunServer(ctx context.Context, pingServer pingservice_proto.PingServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pingservice_proto.RegisterPingServiceServer(server, pingServer)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Info("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Info("Start Ping service port " + port + " ...")
	return server.Serve(listen)
}
