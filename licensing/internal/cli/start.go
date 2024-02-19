package cli

import (
    "github.com/spf13/cobra"
    "github.com/ruby-network/corlink/licensing/internal/server"
    "github.com/ruby-network/corlink/licensing/internal/config"
    "github.com/ruby-network/corlink/licensing/internal/db"
)

func init() {
    rootCmd.AddCommand(startCmd)
    startCmd.Flags().StringP("port", "p", "8080", "Port to listen on")
    startCmd.Flags().StringP("directory", "d", "/", "Directory to serve from")
    startCmd.Flags().StringP("host", "H", "0.0.0.0", "Host to listen on")
    startCmd.Flags().BoolP("sqlite", "s", false, "Use sqlite instead of postgres")
}

var startCmd = &cobra.Command{
    Use:   "start",
    Short: "Start the server",
    Long: `Start the server`,
    Run: func(cmd *cobra.Command, args []string) { 
        host := cmd.Flag("host").Value.String()
        port := cmd.Flag("port").Value.String()
        dir := cmd.Flag("directory").Value.String()
        sqlite, _ := cmd.Flags().GetBool("sqlite")
        config.Init()
        db := db.Init(sqlite)
        server.Start(dir, port, host, db)
    },
}
