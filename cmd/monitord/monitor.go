package main

import (
    "os"
    "log"
    "github.com/spf13/cobra"
)

func main() {
    Cmd := &cobra.Command{
        Use:           "monitord [OPTIONS]",
        Short:         "A self-sufficient runtime for monitors.",
        SilenceUsage:  true,
        SilenceErrors: true,
        RunE: func(cmd *cobra.Command, args []string) error {
            return nil
        },
        Run: func(cmd *cobra.Command, args []string) {
            log.Printf("Inside rootCmd Run with args: %v\n", args)
        },
    }
    
    Cmd.Flags()
    //Flags := Cmd.Flags()
    //Flags.StringVar()
    
    if err := Cmd.Execute(); err != nil {
        log.Printf("cmd error: %v", err)
        os.Exit(1)
    }
}