package command

import (
    "monitor/command/common"
    "github.com/spf13/cobra"
)

var ViewCmd = &common.Command{
    Command:&cobra.Command{
        Use:    common.CMD_VIEW,
        Short:  "view monitor linux info",
        Long:   "View monitor linux info",
        RunE:   func(cmd *cobra.Command, args []string) error {
            
            return nil
        },
    },
}