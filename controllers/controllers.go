package controllers

import (
    "encoding/json"

    "github.com/censhin/go-rest-api/clients"
    "github.com/rackspace/gophercloud/pagination"
    "github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
)

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

func NovaListController() ([]byte, error) {
    client := rackspace.GetClient()

    pager := servers.List(client, nil)

    m := make(map[string][]map[string]interface{})

    err := pager.EachPage(func(page pagination.Page) (bool, error) {
        serverList, err := servers.ExtractServers(page)

        for _, s := range serverList {
            m = append(m, s)
        }
        return true, err
    })

    return nil, err
}
