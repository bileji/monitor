package main

import (
    "os"
    "fmt"
    "github.com/spf13/cobra"
)

func NewCommand(RootCmd *cobra.Command) *cobra.Command {
    HelpCommand := &cobra.Command{
        Use: "help [OPTION]",
        Short: "show help message",
    }
    
    RootCmd.AddCommand(HelpCommand)
    return RootCmd
}

func main() {
    
    RootCmd := &cobra.Command{
        Use: "monitord [OPTIONS]",
        Short: "Linux server status monitor",
        Run: func(cmd *cobra.Command, args []string) {
            // TODO ...
        },
    }
    
    // 添加子命令
    NewCommand(RootCmd)
    
    // TODO flags
    //Flags := RootCmd.Flags()
    //Flags.StringVar(&version, "version", "1.0.0", "this monitor version")
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Printf("%v", err)
        os.Exit(1)
    }
    
}