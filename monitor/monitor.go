package monitor

import (
    "fmt"
    "strconv"
    "gopkg.in/mgo.v2"
    "monitor/monitor/server"
    "monitor/monitor/helper"
    "encoding/json"
    "errors"
    "monitor/monitor/header"
)

const (
    // 定义角色
    NAN int = 0
    MAN int = 1
    NOD int = 2
)

type Monitor struct {
    WebRole *helper.UniqueID
}

type ServerC struct {
    Addr     string            `json:"addr"`
    Database header.Database   `json:"database"`
    Token    string            `json:"token"`
}

func (m *Monitor) SInit(Msg []byte) error {
    Conf := header.Manager{}
    json.Unmarshal(Msg, &Conf)
    
    fmt.Println(Conf)
    
    switch m.WebRole.Get() {
    case MAN:
        return errors.New("has been initialized as manager")
    case NOD:
        return errors.New("has been initialized as node")
    case NAN:
        DB := Conf.Database
        if S, err := mgo.Dial(DB.Host + ":" + strconv.Itoa(int(DB.Port))); err != nil {
            return err
        } else {
            if err := S.DB(DB.Auth).Login(DB.Username, DB.Password); err != nil {
                return err
            }
            go (&service.Master{
                Addr: Conf.Addr,
                DBHandler: S.DB(DB.Auth),
                Token: Conf.Token,
            }).Listen()
            m.WebRole.Set(MAN)
            return nil
        }
    default:
        return nil
    }
}