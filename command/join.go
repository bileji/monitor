package command

import (
    "monitor/command/common"
    "github.com/spf13/cobra"
    "monitor/monitor/header"
)

var JoinCmd = &common.Command{
    Command:&cobra.Command{
        Use: common.CMD_JOIN,
        Short: "monitor node",
        Long: "monitor node",
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
                Command: common.CMD_JOIN,
                Body: []byte(""),
            })
            Socket.EchoReceive()
            return nil
        },
    },
}

func init() {
    //Flags := JoinCmd.Command.Flags()
    //Flags.StringVarP(&JoinCmd.Flags.Join.Addr, "addr", "", "", "manager addr")
    //Flags.StringVarP(&JoinCmd.Flags.Join.Token, "token", "", "", "join token")
    //
}