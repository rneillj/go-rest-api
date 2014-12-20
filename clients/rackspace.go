package rackspace

import (
    "log"

    "github.com/censhin/go-rest-api/util"

    "github.com/rackspace/gophercloud"
    "github.com/rackspace/gophercloud/rackspace"
)

func BuildClient() *gophercloud.ServiceClient {
    creds := config.GetConfig().Auth.Rackspace

    options := gophercloud.AuthOptions {
        Username: creds.Username,
        APIKey: creds.ApiKey,
        IdentityEndpoint: creds.Endpoint,
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
