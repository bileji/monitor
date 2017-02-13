package header

import (
    "gopkg.in/mgo.v2"
)

type Manager struct {
    Addr     string   `json:"addr"`
    Token    string   `json:"token"`
    Log      bool     `json:"log"`
    Database Database `json:"db_handler"`
    Handler  *mgo.Database
}
