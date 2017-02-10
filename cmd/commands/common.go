package commands

import (
    "github.com/spf13/cobra"
)

//type flags struct {
//
//}

type Command struct {
    Flags    *flags
    Command  *cobra.Command
    Children []*cobra.Command
}

func (c *Command) NewChildren() {
    c.Command.AddCommand(c.Children)
}
