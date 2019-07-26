package core

import "time"

//PingCore interface
type PingCore interface {
	Ping(timestamp int64) int64
}

//PingCoreImplement implement PingCore
type PingCoreImplement struct {
}

//NewPingCore create ping core
func NewPingCore() (PingCore, error) {
	return &PingCoreImplement{}, nil
}

func (core *PingCoreImplement) getTimestame() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//Ping ping api
func (core *PingCoreImplement) Ping(timestamp int64) int64 {
	return core.getTimestame()
}
