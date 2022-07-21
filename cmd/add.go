package cmd

import (
	"fmt"
	"strings"
	"time"

	bullettime "biglebowski.nl/bullettime/pkg"
	"github.com/spf13/cobra"
)

var (
	dStr   string
	tStr   string
	addCmd = &cobra.Command{
		Use:     "add description",
		Aliases: []string{"a"},
		Short:   "Add a bullet",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := time.ParseInLocation("2006-01-02", dStr, time.Local)
			if err != nil {
				return fmt.Errorf("could not parse date string: %s", dStr)
			}
			if tStr != "" {
				t, err := time.Parse("15:04", tStr)
				if err != nil {
					return fmt.Errorf("could not parse time string: %s", t)
				}
				dur, _ := time.ParseDuration(fmt.Sprintf("%vh%vm", t.Hour(), t.Minute()))
				d = d.Add(dur)
			}
			b := bullettime.Bullet{}
			b.Description = strings.Join(args, " ")
			b.DateTime = d
			journal.Add(b)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&dStr, "date", "d", time.Now().Format("2006-01-02"), "Set bullet start date `[yyyy-MM-dd]`")
	addCmd.Flags().StringVarP(&tStr, "time", "t", "", "Set bullet meeting start time `[hh:mm]`")
}
