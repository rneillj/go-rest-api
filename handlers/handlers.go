package handlers

import (
    "net/http"
    "log"

    "github.com/censhin/go-rest-api/controllers"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
    b, err := controllers.TestController()

    if err != nil {
        log.Fatal(err)
    } else {
        w.Write(b)
    }
}

func novaListHandler(w http.ResponseWriter, r *http.Request) {
    b, err := controllers.NovaListController()

    if err != nil {
        log.Fatal(err)
    } else {
        w.Write(b)
    }
}

func InitHandlers() {
    http.HandleFunc("/test", testHandler)
    http.HandleFunc("/servers", novaListHandler)
}
