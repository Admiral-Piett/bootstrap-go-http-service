package handlers

import (
    "encoding/json"
    "github.com/Admiral-Piett/bootstrap-go-http-service/app/models"
    "net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(models.BaseResponse{Code: "NOT_IMPLEMENTED", Message: "Sorry, not yet implemented."})
}
