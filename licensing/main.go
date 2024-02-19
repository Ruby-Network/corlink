package main

import (
    "github.com/ruby-network/corlink/licensing/internal/server"
    "github.com/ruby-network/corlink/licensing/internal/config"
)

func main() {
    config.Init()
    server.Start("/", "8080", "0.0.0.0")
}
