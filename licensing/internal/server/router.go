package server

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "encoding/json"
    "errors"
    "strings"
    "gorm.io/gorm"
    database "github.com/ruby-network/corlink/licensing/internal/db"
)

type Response struct {
    Status string `json:"status"`
    Message string `json:"message"`
}


func verifyContentType(w http.ResponseWriter, r *http.Request) error {
    if r.Header.Get("Content-Type") != "application/json" {
        w.WriteHeader(http.StatusUnsupportedMediaType)
        json.NewEncoder(w).Encode(Response{"error", "Unsupported Media Type"})
        err := errors.New("Unsupported Media Type")
        return err
    }
    return nil
}

func generateRoute (w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    key := r.Header.Get("Authorization")
    user := database.GetUserByApiKey(db, key[7:])
    key = strings.ToLower(key[:6]) + key[6:]
    if key != "bearer " + user.ApiKey || key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    w.WriteHeader(http.StatusOK)
    license := database.GenerateKey(db, user)
    json.NewEncoder(w).Encode(Response{"ok", license})
    return
}

func verifyRoute (w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    key := r.Header.Get("Authorization")
    keyToVerify := r.Header.Get("Key")
    //remove the "Bearer " from the
    user := database.GetUserByApiKey(db, key[7:])
    if key != "Bearer " + user.ApiKey || key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    if !database.VerifyKey(db, keyToVerify) {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(Response{"ok", "Authorized"})
    return
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(Response{"ok", "Server is running"})
}

func InitRoutes(dir string, db *gorm.DB) {
    r := chi.NewRouter()
    r.Get(dir, indexRoute)
    r.Post(dir + "generate", func(w http.ResponseWriter, r *http.Request) { generateRoute(w, r, db) })
    r.Post(dir + "verify", func(w http.ResponseWriter, r *http.Request) { verifyRoute(w, r, db) })
    http.Handle(dir, r)
}
