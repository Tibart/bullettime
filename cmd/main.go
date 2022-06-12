package main

import (
	"fmt"

	"biglebowski.nl/bullettime"
)

func main() {
	// TODO: Add arguments

	///t := bullettime.NewBullet()
	// now := time.Now()
	b := bullettime.Bullets{}
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

	b.Load("./test.json")

	fmt.Println(b.String())

	fmt.Println(b.GetSchedule())

	err := b.Save("./test.json")

	fmt.Println(err)

}
