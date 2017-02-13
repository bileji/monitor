package header

type UnixMsg struct {
    Command   string    `json:"command"`
    Timestamp int64     `json:"timestamp"`
    Body      []byte    `json:"body"`
}