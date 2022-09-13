package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	days    uint = 0
	moveCmd      = &cobra.Command{
		Use:     "move bullet-id",
		Aliases: []string{"m"},
		Short:   "Moved bullet to the next day.",
		Args: func(cmd *cobra.Command, args []string) error {
			if !all && len(args) < 1 {
				return fmt.Errorf("bullet id is mandatory when flag -a, --all is not set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if all {
				return journal.Reschedule()
			}

			// Validate args
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("could not convert '%s' to an integer", args[0])
			}
			if id <= 0 {
				return fmt.Errorf("bullet id must be greater than 0")
			}
			if len(args) > 1 {
				d, err := strconv.Atoi(args[1])
				if err != nil {
					return fmt.Errorf("could not convert '%s' to an integer", args[1])
				}

				return journal.Move(id, d)
			}

			// Reschedule bullet
			if err := journal.RescheduleBullet(id); err != nil {
				return fmt.Errorf("could not move bullet to the next day: %s", err.Error())
			}

			return nil
		},
	}
)

func init() {
	moveCmd.Flags().BoolVarP(&all, "all", "a", false, "Apply move to all scheduled journal entries in the past")
	moveCmd.Flags().UintVarP(&days, "days", "d", 0, "Move journal entry forward or backward a number of days")
	rootCmd.AddCommand(moveCmd)
	// TODO add flag for number of days.
}
