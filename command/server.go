package command

import (
    "fmt"
    "encoding/json"
    "monitor/command/common"
    "github.com/spf13/cobra"
    "monitor/monitor/helper"
    "monitor/monitor/header"
)

var InitCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER_INIT,
        Short: "init monitor server",
        Long: "Init monitor server",
        RunE: func(cmd *cobra.Command, args []string) error {
            Socket := &common.Socket{
                SUnix: MainCmd.Configure.Server.Unix,
                CUnix: MainCmd.Configure.Client.Unix,
            }
            err := Socket.UnixSocket();
            if err != nil {
                return err
            }
            defer Socket.Close()
            
            Buffer, _ := json.Marshal(header.Manager{
                Addr: MainCmd.Configure.Server.Addr,
                Database: MainCmd.Configure.MongoDB,
                Token: helper.RandStr(),
            })
            Socket.SendMessage(header.UnixMsg{
                Command: common.CMD_SERVER_INIT,
                Body: Buffer,
            })
            Socket.EchoReceive()
            return nil
        },
    },
}

var TokenCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER_TOKEN,
        Short: "view node join token",
        Long: "View node join token",
        RunE: func(cmd *cobra.Command, args []string) error {
            Socket := &common.Socket{
                SUnix: MainCmd.Configure.Server.Unix,
                CUnix: MainCmd.Configure.Client.Unix,
            }
            err := Socket.UnixSocket();
            if err != nil {
                return err
            }
            defer Socket.Close()
            
            Socket.SendMessage(header.UnixMsg{
                Command: common.CMD_SERVER_TOKEN,
                Body: []byte(""),
            })
            Socket.EchoReceive()
            return nil
        },
    },
}

var ServerCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER,
        Short: "monitor manager",
        Long: "Monitor manager",
        RunE: func(cmd *cobra.Command, args []string) error {
            return nil
        },
    },
}

func init() {
    //Flags := InitCmd.Command.Flags()
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Host, "host", "", "127.0.0.1", "mongodb host")
    //Flags.Int32VarP(&MainCmd.Configure.MongoDB.Port, "port", "", 27017, "mongodb port")
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Auth, "auth", "", "admin", "auth database")
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Username, "user", "", "", "username")
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Password, "pwd", "", "", "password")
    //
    //Flags.StringVarP(&MainCmd.Configure.Server.Addr, "addr", "a", "0.0.0.0:3647", "web server address")
    ServerCmd.NewChildren(InitCmd, TokenCmd)
}