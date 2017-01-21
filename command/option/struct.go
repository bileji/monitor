package option

import (
    "gopkg.in/mgo.v2"
)

type Options struct {
    DBHandler *mgo.Database
}
