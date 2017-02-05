package protocols

type Socket struct {
    Method    string `json:"method"`
    Body      []byte `json:"body"`
    Timestamp int64 `json:"timestamp"`
}