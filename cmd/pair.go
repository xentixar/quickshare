package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"quickshare/helpers"

	"github.com/spf13/cobra"
)

var pairCmd = &cobra.Command{
	Use:  "pair <ip>",
	Long: "This command will send a pairing request to the given ip address",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ipAddress := args[0]
		deviceIPAddress := helpers.GetIPAddress()
		endpoint := fmt.Sprintf("http://%v:8080/pair?ip=%v", ipAddress, deviceIPAddress)
		res, err := http.Get(endpoint)
		if err != nil {
			return err
		}

		if res.StatusCode != 200 {
			body, _ := io.ReadAll(res.Body)
			return errors.New(string(body))
		}

		fmt.Println("Pairing connection send successfully!")
		return nil
	},
}
