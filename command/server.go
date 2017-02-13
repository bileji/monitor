package command

import (
    "encoding/json"
    "monitor/command/common"
    "github.com/spf13/cobra"
    "monitor/monitor"
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
            Buffer, _ := json.Marshal(monitor.ServerC{
                Addr: MainCmd.Configure.Server.Addr,
                Database: MainCmd.Configure.MongoDB,
                Token: helper.RandStr(),
            })
            Socket.SendMessage(header.UnixMsg{
                Command: common.CMD_SERVER_INIT,
                Body: Buffer,
            })
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
    ServerCmd.NewChildren(InitCmd, TokenCmd)
}