package command

import (
    "monitor/command/common"
    "github.com/spf13/cobra"
)

var JoinCmd = &common.Command{
    Command:&cobra.Command{
        Use: common.CMD_JOIN,
        Short: "run as a node",
        Long: "run as a node",
        RunE: func(cmd *cobra.Command, args []string) error {
            // todo
            
            
            return nil
        },
    },
}

func init() {
    
    Flags := JoinCmd.Command.Flags()
    
    Flags.StringVarP(&JoinCmd.Flags.Join.Addr, "addr", "", "", "manager addr")
    Flags.StringVarP(&JoinCmd.Flags.Join.Token, "token", "", "", "join token")
}