package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	cancelCmd = &cobra.Command{
		Use:     "cancel bullet-id",
		Aliases: []string{"cncl", "stop", "rem"},
		Short:   "Cancel bullet.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("could not convert '%s' to an integer", args[0])
			}
			if i <= 0 {
				return fmt.Errorf("bullet id must be greater than 0")
			}

			if err := journal.Cancel(i); err != nil {
				return fmt.Errorf("could cancel bullet: %s", err.Error())
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(cancelCmd)
}
