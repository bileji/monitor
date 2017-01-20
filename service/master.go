package service

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "monitor/collector"
)

type Answer struct {
    Code    int32 `json:"code"`
    Data    map[string]interface{} `json:"data"`
    Message string `json:"message"`
}

func (A *Answer)Return(Res http.ResponseWriter) {
    Res.Header().Set("Content-type", "application/json")
    Bytes, _ := json.Marshal(A)
    Res.Write(Bytes)
}

type Master struct {
    Addr      string
    DBHandler *mgo.Database
}

func (m *Master) Listen() {
    http.HandleFunc("/save", m.Save)
    http.ListenAndServe(m.Addr, nil)
}

func (m *Master) Save(Res http.ResponseWriter, Req *http.Request) {
    var PlayLoad collector.Collector
    
    if Req.Method == "PUT" {
        Body, err := ioutil.ReadAll(Req.Body)
        
        defer Req.Body.Close()
        
        if err != nil {
            (&Answer{
                Code: -1,
                Message: fmt.Sprintf("%v", err),
            }).Return(Res)
            return
        }
        err = json.Unmarshal(Body, &PlayLoad)
        if err != nil {
            (&Answer{
                Code: -1,
                Message: fmt.Sprintf("%v", err),
            }).Return(Res)
            return
        }
        
        // todo...
        (&Answer{
            Code: 0,
            Data: map[string]interface{}{
                "playload": PlayLoad,
            },
            Message: fmt.Sprintf("%v", err),
        }).Return(Res)
        return
    }
    
    (&Answer{
        Code: -1,
        Message: "invalid request",
    }).Return(Res)
}


