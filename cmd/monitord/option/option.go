package option

import (
    "github.com/spf13/pflag"
)

type Common struct {
    Debug bool
}

type Option struct {
    Version    bool
    Common     *Common
    ConfigFile string
    Flags      *pflag.FlagSet
}
