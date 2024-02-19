package cli

import (
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "corlink-server",
    Short: "The licensing server for Corlink",
    Long: `The licensing server for Corlink. This server is used to manage the licensing for the Corlink project.`,
}

func Init() {
    err := rootCmd.Execute()
    if err != nil { os.Exit(1) }
}
