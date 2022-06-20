package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"biglebowski.nl/bullettime"
)

const (
	version    = 0.1
	configPath = "./data.json"
)

var fa string
var fd string
var ft string
var fc int
var fm int
var fp int
var fr int
var fdel int
var fl bool

func init() {
	// Add flags
	flag.StringVar(&fa, "a", "", "Set bullet `description`. Use quotation mark when description has more than one word.")
	flag.StringVar(&fd, "d", time.Now().Format("2006-01-02"), "Set bullet start `date (yyyy-MM-dd)`")
	flag.StringVar(&ft, "t", "", "Set bullet meeting start `time (hh:mm)`")
	// Modifier flags
	flag.IntVar(&fc, "c", -1, "Complete bullet `number` (mandatory")
	flag.IntVar(&fm, "m", -1, "Move bullet `number` to next day")
	flag.IntVar(&fp, "p", -1, "Postpone bullet `number`")
	flag.IntVar(&fr, "r", -1, "Revoke bullet `number`")
	flag.IntVar(&fdel, "del", -1, "delete bullet `number` from schedule")
	// Presentation
	flag.BoolVar(&fl, "l", false, "List all open bullets")

	flag.Usage = func() {
		fmt.Println("Usage of the program:")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	// Construct bullet list
	b := bullettime.Bullets{}
	b.Load(configPath)

	//Set function flags
	if fa != "" {
		if fa[0:1] == "-" {
			fmt.Println("flag needs an argument: -a")
			flag.Usage()
			os.Exit(2)
		}
		if flag.NArg() > 0 {
			fmt.Println("flag argument needs quotation marks: -a")
			flag.Usage()
			os.Exit(2)
		}
		blt := bullettime.Bullet{}
		blt.Description = fa
		d, err := time.ParseInLocation("2006-01-02", fd, time.Local)
		if err != nil {
			log.Printf("could not parse date string: %s\n", fd)
			flag.Usage()
			os.Exit(1)
		}
		blt.DateTime = d
		// Add time when given
		if ft != "" {
			t, err := time.Parse("15:04", ft)
			if err != nil {
				log.Printf("could not parse time string: %s\n", ft)
				os.Exit(1)
			}
			// BUG: Must be better way to get duration
			dur, _ := time.ParseDuration(fmt.Sprintf("%vh%vm", t.Hour(), t.Minute()))
			blt.DateTime = blt.DateTime.Add(dur)
		}
		// Add result
		b.Add(blt)
	}

	// Complete bullet
	if fc > 0 {
		if err := b.Complete(fc); err != nil {
			log.Printf("could not complete bullet: %s", err.Error())
			os.Exit(1)
		}
	}

	// Move bullet to next day
	if fm > 0 {
		if err := b.Reschedule(fm, 1); err != nil {
			log.Printf("could not move bullet to the next day: %s", err.Error())
			os.Exit(1)
		}
	}

	// Postpone bullet
	if fp > 0 {
		if err := b.Postpone(fp); err != nil {
			log.Printf("could not postpone bullet: %s", err.Error())
			os.Exit(1)
		}
	}

	// Cancel bullet
	if fr > 0 {
		if err := b.Cancel(fr); err != nil {
			log.Printf("could not revoke bullet: %s", err.Error())
			os.Exit(1)
		}
	}

	// Delete bullet
	if fdel > 0 {
		if err := b.Remove(fdel); err != nil {
			log.Printf("could not remove bullet from schedule: %s", err.Error())
			os.Exit(1)
		}
	}

	// List bullet journal
	if fl {
		fmt.Println(b.GetSchedule().String())
	}

	b.Save(configPath)
}
