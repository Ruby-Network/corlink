package server

import (
    "net/http"
    "gorm.io/gorm"
    "github.com/fatih/color"
)

func Start(dir string, port string, host string, db *gorm.DB) {
    if dir[0] != '/' { dir = "/" + dir }
    if dir[len(dir) - 1] != '/' { dir = dir + "/" }
    InitRoutes(dir, db)
    color.Green("Server started on http://" + host + ":" + port + dir)
    if host == "0.0.0.0" { color.Blue("You can also access the server at http://localhost:" + port + dir) }
    http.ListenAndServe(host + ":" + port, nil)
}
