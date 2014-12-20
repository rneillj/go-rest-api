package controllers

import (
    "encoding/json"
    "log"

    "github.com/censhin/go-rest-api/resources"
    "github.com/censhin/go-rest-api/util/cache"
)

func TestController() ([]byte, error) {
    body := make(map[string]interface{})

    body["msg"] = "test"

    res, err := json.Marshal(body)

    return res, err
}

func NovaListController() ([]byte) {
    body := make(map[string]interface{})

    res, err := cache.GetValue("listServers")
    if err != nil {
        list, err := resources.NovaListResource()
        if err != nil {
           log.Fatal(err)
        }

        body["servers"] = list

        res, err := json.Marshal(body)
        if err != nil {
            log.Fatal(err)
        }

        cache.CacheValue(res, "listServers")

        return res
    } else {
        return res
    }
}
