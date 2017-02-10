package common

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

type MonitorFlags struct {
    Config string
    Daemon bool
    Pid    string
    Log    string
    Debug  string
}

type ServerFlags struct {
    
}

type JoinFlags struct {
    Addr  string
    Token string
}

type Flags struct {
    Main   MonitorFlags
    Server ServerFlags
    Join   JoinFlags
}

type Command struct {
    Flags     Flags
    Command   *cobra.Command
    Children  []*Command
    Viper     *viper.Viper
    Configure *Configure
}

func (c *Command) NewChildren() {
    for _, Child := range c.Children {
        c.Command.AddCommand(Child.Command)
    }
}

