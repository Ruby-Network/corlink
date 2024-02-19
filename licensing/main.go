package main

import (
    "github.com/ruby-network/corlink/licensing/internal/server"
    "github.com/ruby-network/corlink/licensing/internal/config"
    "github.com/ruby-network/corlink/licensing/internal/db"
)

func main() {
    config.Init()
    db := db.Init()
    server.Start("/", "8080", "0.0.0.0", db)
}
