package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, reg *http.Request) {
			fmt.Fprint(w, "Hello world!")
		})

		// Load port from environment
		port := os.Getenv("PORT")

		// Start the server
		log.Println("HTTP server listening on port: " + port)
		http.ListenAndServe(":"+port, nil)

	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
