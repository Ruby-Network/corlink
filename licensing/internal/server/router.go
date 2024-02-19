package server

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "encoding/json"
    "errors"
)

type Response struct {
    Status string `json:"status"`
    Message string `json:"message"`
}

func verifyContentType(w http.ResponseWriter, r *http.Request) error {
    //only accept application/json 
    if r.Header.Get("Content-Type") != "application/json" {
        w.WriteHeader(http.StatusUnsupportedMediaType)
        json.NewEncoder(w).Encode(Response{"error", "Unsupported Media Type"})
        err := errors.New("Unsupported Media Type")
        return err
    }
    return nil
}

func generateRoute (w http.ResponseWriter, r *http.Request) {
    err := verifyContentType(w, r)
    if err != nil { return }
    key := r.Header.Get("Authorization")
    if key != "Bearer " + "admin" || key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(Response{"ok", "License generated"})
    return
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(Response{"ok", "Server is running"})
}

func InitRoutes(dir string) {
    r := chi.NewRouter()
    r.Get(dir, indexRoute)
    r.Post(dir + "generate", generateRoute)
    http.Handle(dir, r)
}
