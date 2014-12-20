package controllers

import (
    "encoding/json"
    "log"

    "github.com/censhin/go-rest-api/clients"
    "github.com/rackspace/gophercloud/pagination"
    os "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
    rs "github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
)

func TestController() ([]byte, error) {
    body := make(map[string]interface{})

    body["msg"] = "test"

    res, err := json.Marshal(body)

    return res, err
}

func NovaListController() ([]byte, error) {
    client := rackspace.BuildClient()
    pager := rs.List(client, nil)
    data := []os.Server{}
    novaResponse := make(map[string]interface{})

    err := pager.EachPage(func(page pagination.Page) (bool, error) {
        serverList, err := rs.ExtractServers(page)
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
