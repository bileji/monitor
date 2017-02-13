package service

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

type Master struct {
    Addr      string
    DBHandler *mgo.Database
    Token     string
    Log       bool
}

func (m *Master) Listen() {
    http.HandleFunc("/gather", m.Gather)
    http.HandleFunc("/verify", m.Verify)
    err := http.ListenAndServe(m.Addr, nil)
    if err != nil {
        log.Println(err)
    }
}

func (m *Master) Debug(Req *http.Request) {
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

func (m *Master) Gather(Res http.ResponseWriter, Req *http.Request) {
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
        
        err = m.DBHandler.C(collection.GATHER).Insert(Gather);
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

func (m *Master) Verify(Res http.ResponseWriter, Req *http.Request) {
    
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