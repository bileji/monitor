package helper

import (
    "strings"
    "net/http"
)

func Request(Method string, Uri string, Text string) (*http.Response, error) {
    Client := http.Client{}
    
    Req, err := http.NewRequest(Method, Uri, strings.NewReader(Text))
    if err != nil {
        return nil, err
    }
    
    return Client.Do(Req)
}