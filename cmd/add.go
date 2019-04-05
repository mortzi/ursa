package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/mortzi/ursa/data"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a url to be remembered",
	RunE:  addCmdRunE,
}

func addCmdRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("provide url to add")
	}
	url := args[0]
	cat := ""
	tag := ""

	if len(args) >= 2 {
		cat = args[1]
	}
	if len(args) >= 3 {
		tag = args[2]
	}

	resCh, errCh := data.Repository.AddURL(context.Background(), &data.UrsaURL{
		URL:      url,
		Category: cat,
		Tag:      tag,
	})

	for {
		select {
		case res, ok := <-resCh:
			if !ok {
				return nil
			}

			fmt.Printf("%s\n", res)
		case err := <-errCh:
			return err
		}
	}
}

// func addCmdRunE(cmd *cobra.Command, args []string) error {
// 	fileDir := path.Join(os.Getenv("APPDATA"), "ursa")
// 	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
// 		err := os.Mkdir(fileDir, os.ModeDir)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	fileName := path.Join(os.Getenv("appdata"), "ursa", "ursa.json")

// 	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	fmt.Println("file opened scsfly :)")

// 	return nil
// }
