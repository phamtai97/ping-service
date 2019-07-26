package pingservice

import (
	"context"
	pingservice_proto "pingservice/pkg/api"
	"pingservice/pkg/core"
)

//PingService ping service
type PingService struct {
	core core.PingCore
}

//NewPingService new ping service
func NewPingService(core core.PingCore) pingservice_proto.PingServiceServer {
	return &PingService{
		core: core,
	}
}

//Ping api
func (service PingService) Ping(ctx context.Context, msgPing *pingservice_proto.Ping) (*pingservice_proto.Pong, error) {
	timestamp := service.core.Ping(msgPing.Timestamp)

	return &pingservice_proto.Pong{
		Timestamp:   timestamp,
		ServiceName: "Ping Service",
	}, nil
}
