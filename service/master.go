package service

import (
    "fmt"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2"
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
    
    fmt.Println(Req.Method)
    fmt.Println("++++++++")
    fmt.Println(Req.Body)
    
    if Req.Method == "PUT" {
        
    }
}


