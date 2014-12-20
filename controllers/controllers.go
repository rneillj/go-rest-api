package controllers

import (
    "encoding/json"
    "log"

    "github.com/censhin/go-rest-api/clients"
    "github.com/rackspace/gophercloud/pagination"
    os "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
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
    data := []os.Server{}
    novaResponse := make(map[string]interface{})

    err := pager.EachPage(func(page pagination.Page) (bool, error) {
        serverList, err := servers.ExtractServers(page)
        if err != nil {
            return false, err
        }

        for _, s := range serverList {
            data = append(data, s)
        }
        return true, nil
    })

    if err != nil {
        log.Fatal(err)
    }

    novaResponse["servers"] = data

    res, err := json.Marshal(novaResponse)

    return res, err
}
