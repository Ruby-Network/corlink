package server

import (
    "net/http"
    "fmt"
    "gorm.io/gorm"
)

func Start(dir string, port string, host string, db *gorm.DB) {
    if dir[0] != '/' { dir = "/" + dir }
    if dir[len(dir) - 1] != '/' { dir = dir + "/" }
    InitRoutes(dir, db)
    fmt.Println("Server started on http://" + host + ":" + port + dir)
    if host == "0.0.0.0" { fmt.Println("You can access the server at http://localhost:" + port + dir) }
    http.ListenAndServe(host + ":" + port, nil)
}
