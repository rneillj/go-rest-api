package resources

import (
    "github.com/censhin/go-rest-api/clients"

    "github.com/rackspace/gophercloud/pagination"
    os "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
    rs "github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
)

func NovaListResource() ([]os.Server, error) {
    client := rackspace.BuildClient()
    pager := rs.List(client, nil)
    data := []os.Server{}

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

    return data, err
}