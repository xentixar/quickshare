package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"quickshare/http"

	"github.com/spf13/cobra"
)

var shareCmd = &cobra.Command{
	Use:  "share <files>",
	Args: cobra.MinimumNArgs(1),
	Long: "This command accepts the files and create a http server with the given files",
	RunE: func(cmd *cobra.Command, args []string) error {
		files := map[string]string{}

		for _, file := range args {
			info, err := os.Stat(file)
			if err != nil {
				return errors.New("file doesn't exists")
			}

			if info.IsDir() {
				return errors.New("expecting a file, got folder")
			}

			abs, err := filepath.Abs(file)
			if err != nil {
				return errors.New(err.Error())
			}

			files[abs] = file
		}

		http.InitServer(files)
		return nil
	},
}
