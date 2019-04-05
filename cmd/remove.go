package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/mortzi/ursa/data"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "removes a url",
	RunE:    removeCmdRunE,
}

func removeCmdRunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("provide url to be removed")
	}
	resCh, errCh := data.Repository.DeleteURLByURL(context.Background(), args...)

	select {
	case res := <-resCh:
		fmt.Printf("%d urls were removed\n", res)
		return nil
	case err := <-errCh:
		return err
	}
}
