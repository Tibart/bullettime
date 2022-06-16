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

func init() {
	flag.StringVar(&fa, "a", "", "Set bullet `task` description")
	flag.StringVar(&fd, "d", time.Now().Format("2006-01-02"), "Set bullet start `date` (yyyy-MM-dd)")
	flag.StringVar(&ft, "t", "", "Set bullet meeting start `time` (hh:mm)")
	flag.IntVar(&fc, "c", -1, "Complete bullet `number` (mandatory")
	flag.IntVar(&fm, "m", -1, "Move bullet `number` to next day")
	flag.IntVar(&fp, "p", -1, "Postpone bullet `number`")
	flag.IntVar(&fr, "r", -1, "Revoke bullet `number`")
	flag.Usage = func() {
		fmt.Printf("Bullet time, version %v\nUsage of the program: \n", version)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	// Construct bullet list
	b := bullettime.Bullets{}
	b.Load(configPath)

	//Set function flags
	blt := bullettime.Bullet{}
	if fa != "" {
		blt.Description = fa
		d, err := time.ParseInLocation("2006-01-02", fd, time.Local)
		if err != nil {
			log.Printf("could not parse date string: %s\n", fd)
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

	fmt.Println(b.String())

	// Interpreted flags
	// switch strings.ToUpper(os.Args[1]) {
	// case "ADD":
	// 	if *add_b == "" {
	// 		log.Println("flag -b (bullet description) is mandatory!")
	// 		os.Exit(1)
	// 	}
	// 	d, err := time.ParseInLocation("2006-01-02", *add_d, time.Local)
	// 	if err != nil {
	// 		log.Printf("could not parse date string: %s\n", *add_d)
	// 		os.Exit(1)
	// 	}
	// 	if *add_m {
	// 		t, err := time.Parse("15:04", *add_t)
	// 		if err != nil {
	// 			log.Printf("could not parse time string: %s\n", *add_t)
	// 			os.Exit(1)
	// 		}
	// 		// BUG: Must be better way to get duration
	// 		dur, _ := time.ParseDuration(fmt.Sprintf("%vh%vm", t.Hour(), t.Minute()))
	// 		d = d.Add(dur)
	// 	}

	// 	if err = b.Add(*add_b, *add_r, *add_m, d); err != nil {
	// 		log.Println(err.Error())
	// 	}
	// case "LIST":
	// 	fmt.Println(b.String())
	// case "COMPLETE":
	// 	i := *comp_n
	// 	if i < 0 {
	// 		log.Println("flag -n (bullet number) is mandatory!")
	// 		os.Exit(1)
	// 	}
	// 	if i > len(b) {
	// 		log.Println("bullet number does not exist")
	// 		os.Exit(1)
	// 	}

	// 	b.Complete(i)
	// }

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

	//fmt.Println(b.String())

	b.Save(configPath)

	//err := b.Save("./test.json")

	//fmt.Println(err)

}
