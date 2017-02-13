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
    WebRole helper.UniqueID
}

func (m *Monitor) SInit(Msg []byte) error {
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
            return nil
        }
        go Manager.Listen()
        m.WebRole.Set(MAN)
        return nil
    default:
        return nil
    }
}