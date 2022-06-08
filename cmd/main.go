package main

import (
	"fmt"
	"time"

	"biglebowski.nl/bullettime"
)

func main() {
	///t := bullettime.NewBullet()
	b := bullettime.Bullets{}
	b.Add("One", "", time.Now())
	b.Add("Two", "", time.Now())
	b.Add("Three", "", time.Now())
	b.Add("Four", "", time.Now())
	b.Add("Six", "", time.Now())
	b.Add("Seven", "", time.Now())

	b.Remove(2)
	b.SetCompleted(2)
	b.Reschedule(3, 2)
	b.Postpone(4)
	b.Cancel(5)

	fmt.Println(b.String())

}
