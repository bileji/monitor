package command

import (
    "monitor/command/common"
    "github.com/spf13/cobra"
)

var initCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER_INIT,
        Short: "",
        Long: "",
    },
}

var tokenCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER_TOKEN,
        Short: "",
        Long: "",
    },
}

var ServerCmd = &common.Command{
    Command: &cobra.Command{
        Use: common.CMD_SERVER,
        Short: "",
        Long: "",
    },
    Children: []*common.Command{initCmd, tokenCmd},
}