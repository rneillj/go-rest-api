package rackspace

import (
    "fmt"

    "github.com/censhin/go-rest-api/util"

    //"github.com/rackspace/gophercloud"
)

type Config struct {
}

func GetClient() /*(*gophercloud.ProviderClient, error)*/ {
    config := config.GetConfig()
    auth_conf := config["auth"]
    fmt.Println(auth_conf)
    /*
    options := gophercloud.AuthOptions {
        Username: config["auth"]["rackspace"]["username"],
        APIKey: config["auth"]["rackspace"]["apiKey"],
        IdentityEndpoint: config["auth"]["rackspace"]["endpoint"],
    }

    client, err := gophercloud.AuthenticatedClient(options)

    if err != nil {
        return nil, err
    } else {
        return client, nil
    }
    */
}
