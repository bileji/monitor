package commands

import (
    "fmt"
    "runtime"
    "github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
    Use:   "version",
    Short: "print the version number of monitor",
    Long:  `All software has versions. This is monitor's.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        monitorVersion()
        return nil
    },
}

func monitorVersion() {
    fmt.Printf("monitor: v%s %s/%s \n", "1.0.0", runtime.GOOS, runtime.GOARCH)
}