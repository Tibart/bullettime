package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"biglebowski.nl/bullettime"
)

func main() {
	// Capture arguments
	add := flag.NewFlagSet("add", flag.ExitOnError)
	add_bd := add.String("bd", "", "Bullet description (mandatory)")
	add_ref := add.String("ref", "", "External reference id")
	add_d := add.String("d", time.Now().Format("2006-01-02"), "Bullet start date (yyyy-MM-dd)")
	add_t := add.String("t", "", "Bullet meeting start time (hh:mm)")
	add_m := add.Bool("m", false, "Bullet is an meeting")
	add.Parse(os.Args[2:])
	flag.Parse()

	// Construct bullet list
	b := bullettime.Bullets{}

	// Interpreted flags
	switch strings.ToUpper(os.Args[1]) {
	case "ADD":
		if *add_bd == "" {
			fmt.Println("Flag -bd (bullet description) is mandatory!")
			os.Exit(1)
		}
		d, err := time.ParseInLocation("2006-01-02", *add_d, time.Local)
		if err != nil {
			fmt.Printf("could not parse date string: %s\n", *add_d)
			os.Exit(1)
		}
		if *add_m {
			t, err := time.Parse("15:04", *add_t)
			if err != nil {
				fmt.Printf("could not parse time string: %s\n", *add_t)
				os.Exit(1)
			}
			// BUG: Must be better way to get duration
			dur, _ := time.ParseDuration(fmt.Sprintf("%vh%vm", t.Hour(), t.Minute()))
			d = d.Add(dur)
		}

		if err = b.Add(*add_bd, *add_ref, *add_m, d); err != nil {
			fmt.Println(err.Error())
		}

		//fmt.Println(*add_bd, *add_ref, *add_d, *add_t, *add_m)
	}

	//fmt.Println(*add_bd, *add_ref, *add_d, *add_t, *add_m)

	///t := bullettime.NewBullet()
	// now := time.Now()
	//b := bullettime.Bullets{}
	// b.Add("One", "", true, now)
	// b.Add("Two", "", false, now)
	// b.Add("Three", "", false, now)
	// b.Add("Four", "", false, now)
	// b.Add("Six", "", false, now)
	// b.Add("Seven", "", false, now)

	// b.Remove(2)
	// b.Complete(2)
	// b.Reschedule(3, 2)
	// b.Postpone(4)
	// b.Cancel(5)

	//b.Load("./test.json")

	fmt.Println(b.String())

	fmt.Println(b.GetSchedule())

	//err := b.Save("./test.json")

	//fmt.Println(err)

}
