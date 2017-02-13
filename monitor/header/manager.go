package header

type Manager struct {
    Addr     string   `json:"addr"`
    Token    string   `json:"token"`
    Log      bool     `json:"log"`
    Database Database `json:"db_handler"`
}
