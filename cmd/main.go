package main

import (
	"fmt"

	"biglebowski.nl/bullettime"
)

func main() {
	///t := bullettime.NewBullet()
	b := bullettime.Bullets{}
	b.Add("One")
	b.Add("Two")
	b.Add("Three")
	b.Remove(2)

	b.SetCompleted(2)

	fmt.Println(b.String())

}
