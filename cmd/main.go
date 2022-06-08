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

	for _, v := range b {
		fmt.Println(v)
	}

}
