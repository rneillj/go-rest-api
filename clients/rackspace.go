package rackspace

import (
    "log"

    "github.com/censhin/go-rest-api/util"
    "github.com/rackspace/gophercloud"
    "github.com/rackspace/gophercloud/rackspace"
)

func GetClient() *gophercloud.ServiceClient {
    rack_creds := config.GetConfig().Auth.Rackspace

    options := gophercloud.AuthOptions {
        Username: rack_creds.Username,
        APIKey: rack_creds.ApiKey,
        IdentityEndpoint: rack_creds.Endpoint,
    }

    provider, err := rackspace.AuthenticatedClient(options)
    if err != nil {
        log.Fatal(err)
    }

    client, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{Region: "IAD",})
    if err != nil {
        log.Fatal(err)
    }

    return client
}
