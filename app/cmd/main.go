package main

import (
    "fmt"
    "github.com/Admiral-Piett/bootstrap-go-http-service/app"
    "github.com/Admiral-Piett/bootstrap-go-http-service/app/models"
    "log"
    "net/http"
)

func main() {
    service := app.New()

    http.HandleFunc("/", service.ProxyHandler)

    service.Logger.Info("App up and running on http://localhost:", models.SETTINGS.Port)

    // We don't have to directly pass in the router here, since we're assigning the ProxyHandler into the
    //http server itself, and all the routes go through there anyway for CORS.
    err := http.ListenAndServe(fmt.Sprintf(":%s", models.SETTINGS.Port), nil)
    if err != nil {
        log.Fatal(err)
    }
}
