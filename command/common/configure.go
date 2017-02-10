package common

import (
    "fmt"
    "strings"
    "path/filepath"
    "github.com/spf13/viper"
)

type Server struct {
    Daemon bool
    Pid    string
    Unix   string
    Log    string
}

type Client struct {
    UnixFile string
}

type Database struct {
    Host     string `json:"host"`
    Port     int32  `json:"port"`
    Auth     string `json:"auth"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Configure struct {
    Server  Server
    Client  Client
    MongoDB Database
}

func ReadConf(Viper *viper.Viper, Path string) *Configure {
    Dir, File := filepath.Split(Path)
    Ext := filepath.Ext(File)
    
    Viper.SetConfigName(strings.Replace(File, Ext, "", 1))
    Viper.SetConfigType(strings.Trim(Ext, "."))
    Viper.AddConfigPath(Dir)
    
    if Viper.ReadInConfig() != nil {
        fmt.Println("No config file found. Using built-in defaults.")
    }
    
    return &Configure{
        Server: Server{
            Daemon: Viper.GetBool("server.daemon"),
            Pid:  Viper.GetString("server.pid"),
            Log: Viper.GetString("server.log"),
            Unix: Viper.GetString("server.unix"),
        },
        Client: Client{
            UnixFile: Viper.GetString("client.unix"),
        },
        MongoDB: Database{
            Host: Viper.GetString("mongodb.host"),
            Port: int32(Viper.GetInt("mongodb.port")),
            Auth: Viper.GetString("mongodb.auth"),
            Username: Viper.GetString("mongodb.username"),
            Password: Viper.GetString("mongodb.password"),
        },
    }
}

