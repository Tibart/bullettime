package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	completeCmd = &cobra.Command{
		Use:     "complete bullet-id [note]",
		Aliases: []string{"compl", "comp", "cpl", "c"},
		Short:   "Complete bullet.",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("could not convert '%s' to an integer", args[0])
			}
			if i <= 0 {
				return fmt.Errorf("bullet id must be greater than 0")
			}
			note := strings.Join(args[1:], " ")

			if err := journal.Complete(i, note); err != nil {
				return fmt.Errorf("could complete bullet: %s", err.Error())
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(completeCmd)
}
