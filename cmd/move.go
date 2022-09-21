package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	days      int  = 0
	forceMove bool = false
	moveCmd        = &cobra.Command{
		Use:     "move bullet-id",
		Aliases: []string{"m"},
		Short:   "Reschedule (move) bullet to given dasy.",
		Args: func(cmd *cobra.Command, args []string) error {
			if !all && len(args) < 1 {
				return fmt.Errorf("bullet id is mandatory when flag -a, --all is not set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Catch all flag
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

			// Reschedule bullet forward
			if days > 0 {
				// Catch force move flag
				if forceMove {
					return journal.Move(id, days)
				}

				return journal.RescheduleBullet(id, days)
			}

			// Reschedule bullet to today
			return journal.RescheduleBullet(id, 0)
		},
	}
)

func init() {
	moveCmd.Flags().BoolVarP(&all, "all", "a", false, "Move all scheduled journal entries from the past to today")
	moveCmd.Flags().IntVarP(&days, "days", "d", 0, "Move journal entry forward a number of days")
	moveCmd.Flags().BoolVarP(&forceMove, "force", "f", false, "Move bullet forward or backward a couple of days without rescheduling")
	rootCmd.AddCommand(moveCmd)
}
