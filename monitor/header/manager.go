package header

import (
    "gopkg.in/mgo.v2"
)

type ManagerRec struct {
    Addr      string
    Token     string
    Log       bool
    DBHandler *mgo.Database
}
