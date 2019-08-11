package benchmark

import (
	"context"
	"fmt"
	"pingservice/configs"
	pingservice_proto "pingservice/pkg/api"
	"time"

	grpcpool "github.com/processout/grpc-go-pool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//ManagerClient mansger client struct
type ManagerClient struct {
	pool   *grpcpool.Pool
	config *configs.PingServiceConfig
}

//ClientPing client ping struct
type ClientPing struct {
	client pingservice_proto.PingServiceClient
	conn   *grpcpool.ClientConn
	ctx    context.Context
}

//NewManagerClient creat manager client
func NewManagerClient(config *configs.PingServiceConfig) *ManagerClient {
	manager := &ManagerClient{
		config: config,
	}

	p, err := grpcpool.New(manager.NewFactoryClient, config.PoolSize, config.PoolSize, time.Duration(config.TimeOut)*time.Second)
	if err != nil {
		log.Fatal("Do not init connection pool")
	}

	manager.pool = p

	return manager
}

//NewFactoryClient create factory client
func (manager *ManagerClient) NewFactoryClient() (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", manager.config.GRPCHost, manager.config.GRPCPort)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Fatal("Did not connect server")
		return nil, err
	}

	return conn, nil
}

//NewClient new client
func (manager *ManagerClient) NewClient() *ClientPing {
	ctx := context.Background()

	conn, _ := manager.pool.Get(ctx)
	return &ClientPing{
		client: pingservice_proto.NewPingServiceClient(conn.ClientConn),
		conn:   conn,
		ctx:    ctx,
	}
}

//ClosePool close pool
func (manager *ManagerClient) ClosePool() {
	manager.pool.Close()
}

func (c *ClientPing) getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//Ping ping api
func (c *ClientPing) Ping() (*pingservice_proto.Pong, error) {
	msgPing := &pingservice_proto.Ping{
		Timestamp: c.getTimestamp(),
	}

	return c.client.Ping(c.ctx, msgPing)
}

//Close close conn
func (c *ClientPing) Close() error {
	return c.conn.Close()
}
