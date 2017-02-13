package command

import (
    "encoding/json"
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
            
            Buffer, _ := json.Marshal(header.Node{
                Addr: MainCmd.Flags.Join.Addr,
                Token: MainCmd.Flags.Join.Token,
            })
            Socket.SendMessage(header.UnixMsg{
                Command: common.CMD_JOIN,
                Body: Buffer,
            })
            Socket.EchoReceive()
            return nil
        },
    },
}

func init() {
    Flags := JoinCmd.Command.Flags()
    Flags.StringVarP(&MainCmd.Flags.Join.Addr, "addr", "", "", "manager addr")
    Flags.StringVarP(&MainCmd.Flags.Join.Token, "token", "", "", "join token")
}