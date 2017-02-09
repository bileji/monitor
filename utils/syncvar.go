package utils

import (
    "sync"
)

type SyncVar struct {
    sync.RWMutex
    Role int
}

func (v *SyncVar) Get() int {
    v.Lock()
    defer v.Unlock()
    return v.Role
}

func (v *SyncVar) Set(Role int) {
    v.Lock()
    defer v.Unlock()
    v.Role = Role
}
