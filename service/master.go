package service

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "monitor/collector"
    "monitor/static"
    "monitor/collector/model"
)

type Answer struct {
    Code    int32 `json:"code"`
    Data    map[string]interface{} `json:"data"`
    Message string `json:"message"`
}

func (A Answer)Return(Res http.ResponseWriter) {
    Res.Header().Set("Content-type", "application/json")
    Bytes, _ := json.Marshal(A)
    Res.Write(Bytes)
}

type Master struct {
    Addr      string
    DBHandler *mgo.Database
}

func (m *Master) Listen() {
    http.HandleFunc("/gather", m.Gather)
    http.ListenAndServe(m.Addr, nil)
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
        } else {
            err = json.Unmarshal(Body, &Gather);
            if err != nil {
                Answer{
                    Code: -1,
                    Message: fmt.Sprintf("%v", err),
                }.Return(Res)
                return
            }
            
            err = m.DBHandler.C(static.GATHER).Insert(Gather);
            if err != nil {
                Answer{
                    Code: -1,
                    Message: fmt.Sprintf("%v", err),
                }.Return(Res)
                return
            }
            
            Answer{
                Code: 0,
                Message: "gather success",
            }.Return(Res)
            return
        }
    }
    
    Answer{
        Code: -1,
        Message: "invalid request",
    }.Return(Res)
    return
}


