package commands

import "github.com/spf13/cobra"

// Join's flags
type flags struct {
    Addr  string
    Token string
}

var JoinCmd = &Command{
    Command:&cobra.Command{
        Use:    "join",
        Short:  "",
        Long:   "",
        RunE:   func(cmd *cobra.Command, args []string) error {
            // todo
            
            
            return nil
        },
    },
}

func init() {
    
    Flags := JoinCmd.Command.Flags()
    
    Flags.StringVarP(&JoinCmd.Flags.Addr, "addr", "", "", "manager addr")
    Flags.StringVarP(&JoinCmd.Flags.Token, "token", "", "", "join token")
}