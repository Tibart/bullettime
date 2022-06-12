package main

import (
	"fmt"
	"time"

	"biglebowski.nl/bullettime"
)

func main() {
	///t := bullettime.NewBullet()
	now := time.Now()
	b := bullettime.Bullets{}
	b.Add("One", "", true, now)
	b.Add("Two", "", false, now)
	b.Add("Three", "", false, now)
	b.Add("Four", "", false, now)
	b.Add("Six", "", false, now)
	b.Add("Seven", "", false, now)

	b.Remove(2)
	b.Complete(2)
	b.Reschedule(3, 2)
	b.Postpone(4)
	b.Cancel(5)

	fmt.Println(b.String())

	fmt.Println(b.GetSchedule())

}
