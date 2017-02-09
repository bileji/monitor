package utils

import (
    "net"
    "fmt"
    "time"
    "strconv"
    "encoding/json"
    "crypto/md5"
    "encoding/hex"
    "monitor/cmd/protocols"
)

func UnixTime() int64 {
    return time.Now().Unix()
}

func ParseOutPut(Conn *net.UnixConn) {
    Buffer := make([]byte, protocols.READ_LENGTH)
    Len, err := Conn.Read(Buffer)
    if err == nil {
        Message := protocols.OutPut{}
        json.Unmarshal(Buffer[0:Len], &Message)
        fmt.Printf("%s\n\n  ", "Monitor say:")
        fmt.Println(string(Message.Body))
        fmt.Printf("%s\n", "")
    }
    Conn.Close()
}

func Md5(Str string) string {
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(Str))
    cipherStr := md5Ctx.Sum(nil)
    return hex.EncodeToString(cipherStr)
}

func RandStr() string {
    UnixTime := int(UnixTime())
    return Md5(strconv.Itoa(UnixTime))
}

func UsageTemplate() string {
    return `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "[flags]"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}