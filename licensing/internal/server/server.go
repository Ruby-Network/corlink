package server

import (
    "net/http"
    "fmt"
)

func Start(dir string, port string, host string) {
    if dir[0] != '/' { dir = "/" + dir }
    if dir[len(dir) - 1] != '/' { dir = dir + "/" }
    InitRoutes(dir)
    fmt.Println("Server started on http://" + host + ":" + port)
    if host == "0.0.0.0" { fmt.Println("You can access the server at http://localhost:" + port) }
    http.ListenAndServe(host + ":" + port, nil)
}
