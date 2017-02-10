package command

import (
    "fmt"
    "runtime"
    "monitor/command/common"
    "github.com/spf13/cobra"
)

var VersionCmd = &common.Command{
    Command:&cobra.Command{
        Use:    common.CMD_VERSION,
        Short:  "view monitor version",
        Long:   "All software has versions. This is monitor's",
        RunE:   func(cmd *cobra.Command, args []string) error {
            fmt.Printf("version: v%s %s/%s \n", "1.0.0", runtime.GOOS, runtime.GOARCH)
            return nil
        },
    },
}