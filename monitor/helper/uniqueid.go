package helper

import (
    "sync"
)

type UniqueID struct {
    sync.RWMutex
    ID int
}

func (u *UniqueID) Get() int {
    u.Lock()
    defer u.Unlock()
    return u.ID
}

func (u *UniqueID) Set(number int) {
    u.Lock()
    defer u.Unlock()
    u.ID = number
}