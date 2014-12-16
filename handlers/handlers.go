package handlers

import (
    "net/http"
    "encoding/json"
    "log"
)

type TestJson struct {
    Message string
}

func testHandler(w http.ResponseWriter, r *http.Request) {
    t := &TestJson{"test"}
    resp := make(map[string]interface{})

    resp["msg"] = t.Message

    b, err := json.Marshal(resp)

    if err != nil {
        log.Fatal(err)
    } else {
        w.Write(b)
    }
}

func InitHandlers() {
    http.HandleFunc("/test", testHandler)
}
