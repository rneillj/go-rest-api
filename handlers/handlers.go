package handlers

import (
    "net/http"
    "log"

    "github.com/censhin/go-rest-api/controllers"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
    res, err := controllers.TestController()

    if err != nil {
        log.Fatal(err)
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.Write(res)
    }
}

func novaListHandler(w http.ResponseWriter, r *http.Request) {
    res, err := controllers.NovaListController()

    if err != nil {
        log.Fatal(err)
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.Write(res)
    }
}

func InitHandlers() {
    http.HandleFunc("/test", testHandler)
    http.HandleFunc("/servers", novaListHandler)
}
