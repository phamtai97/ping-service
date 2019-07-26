package benchmark

import (
	"github.com/myzhan/boomer"
	log "github.com/sirupsen/logrus"
)

const PING = "ping"

var managerClient *ManagerClient

//BoomerClient boomer client
type BoomerClient struct {
}

//LoadManagerClient load manager client
func (boomerClient *BoomerClient) LoadManagerClient(manager *ManagerClient) {
	managerClient = manager
}

//Ping ping
func (boomerClient *BoomerClient) Ping() {
	client := managerClient.NewClient()
	start := boomer.Now()
	pong, err := client.Ping()
	elapsed := boomer.Now() - start

	client.Close()

	if err != nil {
		log.Error("Ping Account error: ", err)
		boomer.RecordFailure("tcp", "Ping error", elapsed, err.Error())
	} else {
		boomer.RecordSuccess("tcp", "Ping", elapsed, int64(pong.XXX_Size()))
	}
}

//LoadTask load task
func (boomerClient *BoomerClient) LoadTask(nameTask string, weight int) ([]*boomer.Task, error) {
	taskList := make([]*boomer.Task, 0)

	taskPing := boomerClient.createTask(nameTask, weight)
	taskList = append(taskList, taskPing)

	return taskList, nil
}

//getFuncTask get function task
func (boomerClient *BoomerClient) getFuncTask(nameTask string) func() {
	switch nameTask {
	case PING:
		return boomerClient.Ping
	}

	return nil
}

//createTask create task
func (boomerClient *BoomerClient) createTask(nameTask string, weight int) *boomer.Task {
	return &boomer.Task{
		Name:   nameTask,
		Weight: weight,
		Fn:     boomerClient.getFuncTask(nameTask),
	}
}

//RunTask run task
func (boomerClient *BoomerClient) RunTask(tasks []*boomer.Task) {
	boomer.Run(tasks...)
}
