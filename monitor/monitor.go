package monitor

import (
    "monitor/monitor/server"
    "monitor/monitor/helper"
    "encoding/json"
    "errors"
    "strings"
    "monitor/monitor/collector"
)

const (
    // 定义角色
    NAN int = 0
    MAN int = 1
    NOD int = 2
)

type Monitor struct {
    Addr    string
    Token   string
    WebRole helper.UniqueID
}

func (m *Monitor) Role() string {
    switch m.WebRole.Get() {
    case MAN:
        return "manager"
    case NOD:
        return "node"
    case NAN:
        return "not yet initialized"
    default:
        return "the monitor is misleading"
    }
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
            return errors.New("listen tcp " + Manager.Addr + " address already in use")
        }
        m.Addr = Manager.Addr
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
        if len(m.Addr) <= 0 {
            m.Addr = ":80"
        }
        Addr := strings.Split(m.Addr, ":")
        if len(Addr) != 2 {
            return "monitor join --token " + m.Token
        }
        if len(Addr[0]) == 0 {
            Addr[0] = collector.Network{}.GetPublicIP()
        }
        
        return "monitor join --addr " + strings.Join(Addr, ":") + " --token " + m.Token, nil
    case NOD:
        return "", errors.New("run as a node")
    case NAN:
        return "", errors.New("not yet initialized")
    default:
        return "", errors.New("the monitor is misleading")
    }
}

func (m *Monitor) Join(Msg []byte) error {
    Node := server.Node{}
    json.Unmarshal(Msg, &Node)
    
    return Node.RunForever()
}