package cmd

import (
	"quickshare/http"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:  "start",
	Long: "This command will run the daemon on background",
	RunE: func(cmd *cobra.Command, args []string) error {
		http.StartHTTPDaemon()
		return nil
	},
}
