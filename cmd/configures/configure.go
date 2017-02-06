package configures

import (
    "fmt"
    "strings"
    "path/filepath"
    "github.com/spf13/viper"
)

type Config struct {
    PidFile    string
    DSockFile  string
    CSockFile  string
    LogFile    string

    DBHost     string
    DBPort     int
    DB         string
    DBUsername string
    DBPassword string
}

func Initialize(Viper *viper.Viper, Path string) *viper.Viper {
    Dir, File := filepath.Split(Path)
    Ext := filepath.Ext(File)
    
    Viper.SetConfigName(strings.Replace(File, Ext, "", 1))
    Viper.SetConfigType(strings.Trim(Ext, "."))
    Viper.AddConfigPath(Dir)
    
    if Viper.ReadInConfig() != nil {
        fmt.Println("No config file found. Using built-in defaults.")
    }
    
    return Viper
}
