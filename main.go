package main

import "fmt"

func main() {
	// inisilisasi map
	colors := make(map[string]string)
	// colors := map[string]string{
	// 	"red":    "#fd977",
	// 	"yellow": "#87fd98",
	// }

	// assignment value kedalam map
	colors["white"] = "#FFFFFF"
	colors["black"] = "#000000"

	fmt.Println(colors)
}
