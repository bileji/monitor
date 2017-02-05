package command

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "monitord",
    Short: "Linux server status monitor",
    Long: "",
    Run: func(cmd *cobra.Command, args []string) {
        // TODO ...
    },
}

func addCommand(RootCmd *cobra.Command) {
    RootCmd.AddCommand(serverCmd)
}

func Execute() {
    addCommand(rootCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Printf("%v", err)
        os.Exit(1)
    }
}

// todo add flags
func init() {
    
}
