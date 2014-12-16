package controllers

import "encoding/json"

type TestJson struct {
    Message string
}

func TestController() ([]byte, error) {
    t := &TestJson{"test"}
    body := make(map[string]interface{})

    body["msg"] = t.Message

    b, err := json.Marshal(body)

    return b, err
}
