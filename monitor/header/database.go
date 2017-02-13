package header

type Database struct {
    Host     string `json:"host"`
    Port     int32  `json:"port"`
    Auth     string `json:"auth"`
    Username string `json:"username"`
    Password string `json:"password"`
}