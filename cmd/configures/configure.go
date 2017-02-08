package configures

import (
    "fmt"
    "strings"
    "path/filepath"
    "github.com/spf13/viper"
)

type Server struct {
    Daemon   bool
    PidFile  string
    UnixFile string
    LogFile  string
}

type Client struct {
    UnixFile string
}

type Database struct {
    Host     string `json:"host"`
    Port     int32  `json:"port"`
    AuthDB   string `json:"auth_db"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Conf struct {
    Server  Server
    Client  Client
    MongoDB Database
}

func Initialize(Viper *viper.Viper, Path string) *Conf {
    Dir, File := filepath.Split(Path)
    Ext := filepath.Ext(File)
    
    Viper.SetConfigName(strings.Replace(File, Ext, "", 1))
    Viper.SetConfigType(strings.Trim(Ext, "."))
    Viper.AddConfigPath(Dir)
    
    if Viper.ReadInConfig() != nil {
        fmt.Println("No config file found. Using built-in defaults.")
    }
    
    return &Conf{
        Server: Server{
            Daemon: Viper.GetBool("server.daemon"),
            PidFile:  Viper.GetString("server.pid_file"),
            UnixFile: Viper.GetString("server.unix_file"),
            LogFile: Viper.GetString("server.log_file"),
        },
        Client: Client{
            UnixFile: Viper.GetString("client.unix_file"),
        },
        MongoDB: Database{
            Host: Viper.GetString("mongodb.host"),
            Port: int32(Viper.GetInt("mongodb.port")),
            AuthDB: Viper.GetString("mongodb.auth_db"),
            Username: Viper.GetString("mongodb.username"),
            Password: Viper.GetString("mongodb.password"),
        },
    }
}
