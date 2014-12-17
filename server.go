package main

import (
    //"fmt"
    "net/http"

    "github.com/censhin/go-rest-api/clients"
    "github.com/censhin/go-rest-api/handlers"
)

func main() {
    /*rack_client, err := */rackspace.GetClient()
    //fmt.Println(rack_client)
    handlers.InitHandlers()
    http.ListenAndServe(":8080", nil)
}
