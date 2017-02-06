package configures

import "github.com/spf13/viper"

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

func (c *Config) Initialize(Path string) *viper.Viper {
    Conf := viper.New()
    
    //Conf.SetConfigFile()
    //Conf.AddConfigPath(Path)
    //Conf.SetConfigType("yaml")
    
    return Conf
}
