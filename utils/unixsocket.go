package utils

import (
    "os"
    "net"
    "monitor/cmd/configures"
)

func UnixSocket(Conf *configures.Conf) (*net.UnixConn, error) {
    os.Remove(Conf.Client.UnixFile)
    
    lAddr, err := net.ResolveUnixAddr("unix", Conf.Client.UnixFile)
    if err != nil {
        return nil, err
    }
    
    rAddr, err := net.ResolveUnixAddr("unix", Conf.Server.UnixFile)
    if err != nil {
        return nil, err
    }
    
    return net.DialUnix("unix", lAddr, rAddr)
}
