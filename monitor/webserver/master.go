package service

import (
    "log"
    "fmt"
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
}

func (m *Master) Listen() {
    http.HandleFunc("/gather", m.Gather)
    http.HandleFunc("/verify", m.Verify)
    err := http.ListenAndServe(m.Addr, nil)
    if err != nil {
        log.Println(err)
    }
}

func (m *Master) Gather(Res http.ResponseWriter, Req *http.Request) {
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
    
    log.Println(Req.UserAgent())
    
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
