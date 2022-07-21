package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	postponeCmd = &cobra.Command{
		Use:     "postpone bullet-id",
		Aliases: []string{"postp", "post", "pp"},
		Short:   "Postpone bullet and move it to the backlog.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("could not convert '%s' to an integer", args[0])
			}
			if i <= 0 {
				return fmt.Errorf("bullet id must be greater than 0")
			}

			if err := journal.Postpone(i); err != nil {
				return fmt.Errorf("could postpone bullet: %s", err.Error())
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(postponeCmd)
}
