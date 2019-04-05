package cmd

import (
	"context"
	"fmt"

	"github.com/mortzi/ursa/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "use to list all saved urls",
	RunE:    listCmdRunE,
}

func listCmdRunE(cmd *cobra.Command, args []string) error {
	urlChan, errChan := data.Repository.GetAllURLs(context.Background())
	counter := 0

	for {
		select {
		case url, ok := <-urlChan:
			if !ok {
				return nil
			}

			fmt.Printf("%d=> %s\n", counter, url)
			counter++
		case err := <-errChan:
			return err
		}
	}
}
