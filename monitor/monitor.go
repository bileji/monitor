package monitor

import (
    "monitor/monitor/server"
    "monitor/monitor/helper"
    "encoding/json"
    "errors"
)

const (
    // 定义角色
    NAN int = 0
    MAN int = 1
    NOD int = 2
)

type Monitor struct {
    Token   string
    WebRole helper.UniqueID
}

func (m *Monitor) ManagerInit(Msg []byte) error {
    Manager := server.Manager{}
    json.Unmarshal(Msg, &Manager)
    
    switch m.WebRole.Get() {
    case MAN:
        return errors.New("has been initialized as manager")
    case NOD:
        return errors.New("has been initialized as node")
    case NAN:
        err := Manager.ConnectDB()
        if err != nil {
            return err
        }
        EMsg := make(chan bool, 1)
        go Manager.Listen(EMsg)
        if ! <-EMsg {
            return errors.New("tcp " + Manager.Addr + " address already in use")
        }
        m.Token = Manager.Token
        m.WebRole.Set(MAN)
        return nil
    default:
        return errors.New("the monitor is misleading")
    }
}

func (m *Monitor) ManagerToken() (string, error) {
    switch m.WebRole.Get() {
    case MAN:
        return m.Token, nil
    case NOD:
        return "", errors.New("run as a node")
    case NAN:
        return "", errors.New("not yet initialized")
    default:
        return "", errors.New("the monitor is misleading")
    }
}