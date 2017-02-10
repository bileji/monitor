package protocol

type SocketMsg struct {
    Command   string    `json:"command"`
    Body      []byte    `json:"body"`
    Timestamp int64     `json:"timestamp"`
}
