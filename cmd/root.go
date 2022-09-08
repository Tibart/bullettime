package cmd

import (
	"fmt"
	"log"
	"os"

	bullettime "biglebowski.nl/bullettime/pkg"
	"github.com/spf13/cobra"
)

const (
	// TODO: Checkout viper for config implementation
	version    = "0.1"
	configPath = "./data.json"
)

var (
	all     bool = false
	journal      = &bullettime.Bullets{}
	rootCmd      = &cobra.Command{
		Use:   "bullettime",
		Short: "Execute a Bullet-time command",
		Long: `Bullet-time, a command line interface that helps you to stay focussed 
and keep track of what you planned to do. Also it keeps track 
of what you did of which bullets got postponed or send back to 
the backlog.`,
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			if all {
				fmt.Println(journal.String())
			} else {
				fmt.Println(journal.TodaysSchedule().String())
			}
		},
	}
)

func init() {
	journal.Load(configPath)
	rootCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Show every journal entry.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := journal.Save(configPath); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
