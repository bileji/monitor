package command

import (
    //"encoding/json"
    "monitor/command/common"
    "github.com/spf13/cobra"
    //"monitor/command/protocol"
    //"monitor/monitor"
    //"monitor/monitor/helper"
)

var initCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER_INIT,
        Short: "init monitor server",
        Long: "Init monitor server",
        RunE: func(cmd *cobra.Command, args []string) error {
            //Socket := &common.Socket{
            //    SUnix: MainCmd.Configure.Server.Unix,
            //    CUnix: MainCmd.Configure.Client.Unix,
            //}
            //err := Socket.UnixSocket();
            //if err != nil {
            //    return err
            //}
            //Buffer, _ := json.Marshal(monitor.ServerC{
            //    Addr: MainCmd.Configure.Server.Addr,
            //    Database: MainCmd.Configure.MongoDB,
            //    Token: helper.RandStr(),
            //})
            //Socket.SendMessage(protocol.SocketMsg{
            //    Command: common.CMD_SERVER_INIT,
            //    Body: Buffer,
            //})
            return nil
        },
    },
}

var tokenCmd = &common.Command{
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
    },
    Children: []*common.Command{initCmd, tokenCmd},
}

func init() {
    initFlags := initCmd.Command.Flags()
    initFlags.StringVarP(&MainCmd.Configure.MongoDB.Host, "host", "", "127.0.0.1", "mongodb host")
    initFlags.Int32VarP(&MainCmd.Configure.MongoDB.Port, "port", "", 27017, "mongodb port")
    initFlags.StringVarP(&MainCmd.Configure.MongoDB.Auth, "auth", "", "admin", "auth database")
    initFlags.StringVarP(&MainCmd.Configure.MongoDB.Username, "user", "", "", "username")
    initFlags.StringVarP(&MainCmd.Configure.MongoDB.Password, "pwd", "", "", "password")
    
    initFlags.StringVarP(&MainCmd.Configure.Server.Addr, "addr", "a", "0.0.0.0:3647", "web server address")
}