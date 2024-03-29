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
	version    = "0.3"
	configPath = "./data.json"
)

var (
	all       bool = false
	postponed bool = false
	yesterday bool = false
	week      bool = false
	journal        = &bullettime.Bullets{}
	rootCmd        = &cobra.Command{
		Use:   "bullettime",
		Short: "Execute a Bullet-time command",
		Long: `Bullet-time, a command line interface that helps you to stay focussed 
and keep track of what you planned to do. Also it keeps track 
of what you did of which bullets got postponed or send back to 
the backlog.`,
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			if week {
				fmt.Println("Bullet week journal")
				fmt.Println(journal.Week())
			} else if yesterday {
				fmt.Println("Yesterdays bullets")
				fmt.Println(journal.Yesterday())
			} else if postponed {
				fmt.Println("Postponed bullets")
				fmt.Println(journal.Postponed())
			} else if all {
				fmt.Println("All bullets")
				fmt.Println(journal.String())
			} else {
				fmt.Println("Todays bullets")
				fmt.Println(journal.TodaysSchedule().String())
			}
		},
	}
)

func init() {
	journal.Load(configPath)
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all journal entries")
	rootCmd.Flags().BoolVarP(&postponed, "postponed", "p", false, "Show postponed journal entries")
	rootCmd.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "Show yesterdays journal entries")
	rootCmd.Flags().BoolVarP(&week, "week", "w", false, "Show this weeks journal entries")
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
