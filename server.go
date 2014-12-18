package main

import (
    "net/http"

    "github.com/censhin/go-rest-api/handlers"
)

func main() {
    handlers.InitHandlers()
    http.ListenAndServe(":8080", nil)
}
