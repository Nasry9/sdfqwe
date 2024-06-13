package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Response struct {
    JSONRPC string `json:"jsonrpc"`
    Result  string `json:"result"`
    Error   *Error `json:"error"`
    ID      int    `json:"id"`
}

type Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func getNodeVersion(w http.ResponseWriter, r *http.Request) {
    response := Response{
        JSONRPC: "2.0",
        Result:  "0.1.0",
        Error:   nil,
        ID:      1,
    }
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/ext/info", getNodeVersion)
    log.Println("Starting JSON-RPC server on :8550")
    log.Fatal(http.ListenAndServe(":8550", nil))
}
