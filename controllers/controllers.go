package controllers

import (
    "encoding/json"
    "log"

    "github.com/censhin/go-rest-api/resources"
)

func TestController() ([]byte, error) {
    body := make(map[string]interface{})

    body["msg"] = "test"

    res, err := json.Marshal(body)

    return res, err
}

func NovaListController() ([]byte, error) {
    body := make(map[string]interface{})

    list, err := resources.NovaListResource()
    if err != nil {
        log.Fatal(err)
    } else {
        body["servers"] = list
    }

    res, err := json.Marshal(body)

    return res, err
}
