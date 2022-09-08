package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	moveCmd = &cobra.Command{
		Use:     "move bullet-id",
		Aliases: []string{"m"},
		Short:   "Moved bullet to the next day.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("could not convert '%s' to an integer", args[0])
			}
			if i <= 0 {
				return fmt.Errorf("bullet id must be greater than 0")
			}

			if err := journal.Reschedule(i); err != nil {
				return fmt.Errorf("could not move bullet to the next day: %s", err.Error())
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(moveCmd)
	// TODO add flag for number of days.
}
