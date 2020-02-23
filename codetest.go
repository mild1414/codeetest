package main

import (
	"fmt"
	"os"
)

func getDrives() (r []string) {
	fmt.Print("input : ")
	var get string
	fmt.Scan(&get)
	f, err := os.Open(string(get) + ":\\")
	if err == nil {
		drive := string(get) + ":/"
		r = append(r, drive)
	}
}
