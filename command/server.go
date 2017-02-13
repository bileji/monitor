package command

import (
    "monitor/command/common"
    "github.com/spf13/cobra"
)

var initCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER_INIT,
        Short: "init monitor server",
        Long: "Init monitor server",
        RunE: func(cmd *cobra.Command, args []string) error {
            
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