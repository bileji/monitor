package header

type Answer struct {
    Code    int32                   `json:"code"`
    Data    map[string]interface{}  `json:"data"`
    Message string                  `json:"message"`
}