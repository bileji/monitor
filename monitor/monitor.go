package monitor

import (
    "strconv"
    "monitor/cmd/configures"
    "gopkg.in/mgo.v2"
    "monitor/monitor/webserver"
)

type Monitor struct{}

type WebServer struct {
    Addr     string
    Database configures.Database
}

// 初始化服务
func (m *Monitor) ServerInit(WS *WebServer) error {
    
    Session, err := mgo.Dial(WS.Database.Host + ":" + strconv.Itoa(int(WS.Database.Port)))
    if err != nil {
        return err
    }
    
    err = Session.DB(WS.Database.AuthDB).Login(WS.Database.Username, WS.Database.Password)
    if err != nil {
        return err
    }
    
    go (&service.Master{
        Addr: WS.Addr,
        DBHandler: Session.DB(WS.Database.AuthDB),
    }).Listen()
    
    return nil
}