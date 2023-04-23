package handlers

import (
    "encoding/json"
    "github.com/Admiral-Piett/bootstrap-go-http-service/app/models"
    "net/http"
)

// NotFoundHandler is an example pattern for easy routing to individual functions based on HTTP Method.
//In this case, since there's only one available function we route every request to `notFoundHandler`.
func NotFoundHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        // TODO - set Auth pattern here
        switch r.Method {
        case "GET":
            notFoundHandler(w, r)
        case "POST":
            notFoundHandler(w, r)
        case "PUT":
            notFoundHandler(w, r)
        case "PATCH":
            notFoundHandler(w, r)
        case "DELETE":
            notFoundHandler(w, r)
        }
    }
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(models.BaseResponse{Code: "NOT_FOUND", Message: "Request path not found."})
}
