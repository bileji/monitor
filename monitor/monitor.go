package monitor

import (
    "strconv"
    "gopkg.in/mgo.v2"
    "monitor/monitor/server"
    "monitor/monitor/helper"
    "monitor/command/common"
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
    WebRole *helper.UniqueID
}

type ServerC struct {
    Addr     string            `json:"addr"`
    Database common.Database   `json:"database"`
    Token    string            `json:"token"`
}

func (m *Monitor) SInit(Msg []byte) error {
    var Conf ServerC
    json.Unmarshal(Msg, &Conf)
    
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
            go (&service.Master{Addr: Conf, DBHandler: DB.Auth, Token: Conf.Token}).Listen()
            m.WebRole.Set(MAN)
            return nil
        }
    default:
        return nil
    }
}

//type WebServer struct {
//    Addr     string                 `json:"addr"`
//    Database configures.Database    `json:"database"`
//    Token    string                 `json:"token"`
//}
//
//// 初始化服务
//func (m *Monitor) ServerInit(WS *WebServer) error {
//
//    Session, err := mgo.Dial(WS.Database.Host + ":" + strconv.Itoa(int(WS.Database.Port)))
//    if err != nil {
//        return err
//    }
//
//    err = Session.DB(WS.Database.AuthDB).Login(WS.Database.Username, WS.Database.Password)
//    if err != nil {
//        return err
//    }
//
//    go (&service.Master{
//        Addr: WS.Addr,
//        DBHandler: Session.DB(WS.Database.AuthDB),
//        Token: WS.Token,
//    }).Listen()
//
//    return nil
//}
//
//// node join
//func (m *Monitor) Join() error {
//
//    return nil
//}