package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a url to be remembered",
	RunE:  addCmdRunE,
}

func addCmdRunE(cmd *cobra.Command, args []string) error {
	fileDir := path.Join(os.Getenv("APPDATA"), "ursa")
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err := os.Mkdir(fileDir, os.ModeDir)
		if err != nil {
			return err
		}
	}

	fileName := path.Join(os.Getenv("appdata"), "ursa", "ursa.json")

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("file opened scsfly :)")

	return nil
}
