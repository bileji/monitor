package server

import (
    "log"
    "fmt"
    "strings"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "monitor/monitor/collector/model"
    "monitor/monitor/collector/collection"
    "monitor/monitor/header"
    "strconv"
    "github.com/shirou/gopsutil/net"
)

type Answer struct {
    Code    int32                   `json:"code"`
    Data    map[string]interface{}  `json:"data"`
    Message string                  `json:"message"`
}

func (A Answer)Return(Res http.ResponseWriter) {
    Res.Header().Set("Content-type", "application/json")
    if len(A.Data) <= 0 {
        A.Data = make(map[string]interface{}, 1)
    }
    Bytes, _ := json.Marshal(A)
    Res.Write(Bytes)
}

type Manager header.Manager

func (m *Manager) ConnectDB() error {
    // 创建连接
    Session, err := mgo.Dial(m.Database.Host + ":" + strconv.Itoa(int(m.Database.Port)))
    if err != nil {
        return err
    }
    
    // 登录DB
    if err := Session.DB(m.Database.Auth).Login(m.Database.Username, m.Database.Password); err != nil {
        return err
    }
    
    // 连接池限制
    Session.SetPoolLimit(10)
    m.Handler = Session.DB(m.Database.Auth)
    
    return nil
}

func (m *Manager) CheckPort() error {
    Conn, err := net.Listen("tcp", m.Addr)
    if err == nil {
        Conn.Close()
        return nil
    }
    return err
}

func (m *Manager) Listen() {
    http.HandleFunc("/gather", m.Gather)
    http.HandleFunc("/verify", m.Verify)
    
    err := http.ListenAndServe(m.Addr, nil)
    if err != nil {
        log.Println(err)
    }
}

func (m *Manager) Debug(Req *http.Request) {
    if m.Log == true {
        LogStr := []string{
            "[web]",
            Req.RemoteAddr,
            Req.Method,
            Req.RequestURI,
            Req.UserAgent(),
        }
        log.Println(strings.Join(LogStr, " "))
    }
}

func (m *Manager) Gather(Res http.ResponseWriter, Req *http.Request) {
    m.Debug(Req)
    
    var Gather model.Gather
    
    if Req.Method == "PUT" {
        Body, err := ioutil.ReadAll(Req.Body);
        defer Req.Body.Close()
        
        if err != nil {
            Answer{
                Code: -1,
                Message: fmt.Sprintf("%v", err),
            }.Return(Res)
            return
        }
        
        err = json.Unmarshal(Body, &Gather);
        if err != nil {
            Answer{
                Code: -1,
                Message: fmt.Sprintf("%v", err),
            }.Return(Res)
            return
        }
        
        err = m.Handler.C(collection.GATHER).Insert(Gather);
        if err != nil {
            Answer{
                Code: -1,
                Message: fmt.Sprintf("%v", err),
            }.Return(Res)
            return
        }
        
        Answer{
            Code: 0,
            Message: "gather successful",
        }.Return(Res)
        return
    }
    
    Answer{
        Code: -1,
        Message: "invalid request",
    }.Return(Res)
    return
}

func (m *Manager) Verify(Res http.ResponseWriter, Req *http.Request) {
    
    m.Debug(Req)
    
    if Req.Method == "PUT" {
        Body, err := ioutil.ReadAll(Req.Body)
        defer Req.Body.Close()
        
        if err != nil {
            Answer{
                Code: -1,
                Message: fmt.Sprintf("%v", err),
            }.Return(Res)
            return
        }
        
        if m.Token == string(Body) {
            Answer{
                Code: 0,
                Message: "verify successful",
            }.Return(Res)
            return
        }
        
        Answer{
            Code: -1,
            Message: "token does not match",
        }.Return(Res)
        return
    }
    
    Answer{
        Code: -1,
        Message: "invalid request",
    }.Return(Res)
    return
}
