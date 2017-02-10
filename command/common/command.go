package common

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

type Command struct {
    Flags     map[string]interface{}
    Command   *cobra.Command
    Children  []*cobra.Command
    Viper     *viper.Viper
    Configure *Configure
}

func (c *Command) NewChildren() {
    for _, Child := range c.Children {
        c.Command.AddCommand(Child)
    }
}

