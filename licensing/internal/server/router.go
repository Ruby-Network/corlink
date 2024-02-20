package server

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "encoding/json"
    "errors"
    "strings"
    "gorm.io/gorm"
    //"fmt"
    database "github.com/ruby-network/corlink/licensing/internal/db"
)

type Response struct {
    Status string `json:"status"`
    Message string `json:"message"`
}

func verifyHeaders(headers []string, r *http.Request) error {
    for _, header := range headers {
        if r.Header.Get(header) == "" {
            return errors.New("Missing header " + header)
        } 
    }
    return nil
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

func createApiKey(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    err = verifyHeaders([]string{"Authorization", "User"}, r)
    if err != nil { 
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
    key := r.Header.Get("Authorization")
    user := r.Header.Get("User") 
    isAdminKey := database.OnlyAdmin(db, key[7:])
    if !isAdminKey {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    if key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    if user == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", "User not provided"})
        return
    }
    err, keyFromDB := database.CreateUser(db, user) 
    if err != nil {
        w.WriteHeader(http.StatusConflict)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(Response{"ok", keyFromDB})
    return
}

func deleteApiKey(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    err = verifyHeaders([]string{"Authorization", "User"}, r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
    key := r.Header.Get("Authorization")
    user := r.Header.Get("User")
    isAdminKey := database.OnlyAdmin(db, key[7:])
    if !isAdminKey {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    if key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    } 
    userDeleted := database.DeleteUser(db, user)
    if !userDeleted {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(Response{"error", "User not found"})
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(Response{"ok", "Deleted"})
    return
}

func getApiKey(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    err = verifyHeaders([]string{"Authorization", "User"}, r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
    key := r.Header.Get("Authorization")
    user := r.Header.Get("User")
    isAdminKey := database.OnlyAdmin(db, key[7:])
    if !isAdminKey {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    if key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    if user == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", "User not provided"})
        return
    }
    keyFromDB, validUser := database.GetApiKey(db, user)
    if !validUser {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(Response{"error", "User not found"})
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(Response{"ok", keyFromDB})
    return
}

func updateApiKey(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    err = verifyHeaders([]string{"Authorization", "User"}, r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
    key := r.Header.Get("Authorization")
    user := r.Header.Get("User")
    isAdminKey := database.OnlyAdmin(db, key[7:])
    if !isAdminKey {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    if key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    }
    w.WriteHeader(http.StatusOK)
    keyFromDB := database.UpdateUserKey(db, user) 
    json.NewEncoder(w).Encode(Response{"ok", keyFromDB})
    return
}

func generateRoute (w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    err = verifyHeaders([]string{"Authorization"}, r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
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

func deleteRoute (w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    err = verifyHeaders([]string{"Authorization", "Key"}, r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
    key := r.Header.Get("Authorization")
    keyToVerify := r.Header.Get("Key")
    user := database.GetUserByApiKey(db, key[7:])
    key = strings.ToLower(key[:6]) + key[6:]
    if key != "bearer " + user.ApiKey || key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(Response{"error", "Unauthorized"})
        return
    } 
    w.WriteHeader(http.StatusOK)
    database.DeleteKey(db, keyToVerify)
    json.NewEncoder(w).Encode(Response{"ok", "Deleted"})
    return
}

func verifyRoute (w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    err := verifyContentType(w, r)
    if err != nil { return }
    err = verifyHeaders([]string{"Authorization", "Key"}, r)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(Response{"error", err.Error()})
        return
    }
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
    r.Post(dir + "delete", func(w http.ResponseWriter, r *http.Request) { deleteRoute(w, r, db) })
    r.Post(dir + "create-user", func(w http.ResponseWriter, r *http.Request) { createApiKey(w, r, db) })
    r.Post(dir + "delete-user", func(w http.ResponseWriter, r *http.Request) { deleteApiKey(w, r, db) })
    r.Post(dir + "get-user", func(w http.ResponseWriter, r *http.Request) { getApiKey(w, r, db) })
    r.Post(dir + "update-user", func(w http.ResponseWriter, r *http.Request) { updateApiKey(w, r, db) })
    http.Handle(dir, r)
}
