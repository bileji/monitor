package command

import (
    "github.com/spf13/cobra"
    "monitor/command/common"
    "monitor/monitor/header"
)

var RoleCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_ROLE,
        Short: "view this monitor indentify",
        Long: "view this monitor indentify",
        RunE: func(cmd *cobra.Command, args []string) error {
            Socket := &common.Socket{
                SUnix: MainCmd.Configure.Server.Unix,
                CUnix: MainCmd.Configure.Client.Unix,
            }
            err := Socket.UnixSocket();
            defer Socket.Conn.Close()
            if err != nil {
                return err
            }
            Socket.SendMessage(header.UnixMsg{
                Command: common.CMD_ROLE,
                Body: []byte(""),
            })
            Socket.EchoReceive()
            return nil
        },
    },
}