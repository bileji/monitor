package commands

import (
    "encoding/json"
    "github.com/spf13/cobra"
    "monitor/cmd/configures"
    "monitor/utils"
    "monitor/cmd/protocols"
)

var RoleCmd = &cobra.Command{
    Use: "role",
    Short: "view roles",
    Long: "View roles",
    RunE: func(cmd *cobra.Command, args []string) error {
        Conf := configures.Initialize(Viper, File.Conf)
        
        Conn, err := utils.UnixSocket(Conf)
        if err != nil {
            return err
        }
        
        Message, _ := json.Marshal(protocols.Socket{
            Command: protocols.ROLE,
            Body: []byte(""),
            Timestamp: utils.UnixTime(),
        })
        Conn.Write(Message)
        
        utils.ParseOutPut(Conn)
        
        return nil
    },
}