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
		d := string(get) + ":/"
		r = append(r, d)
		f.Close()
	}
	return
}

func FindFileFromExtension(extension []string, dir string, files *[]string, Path *[]string) {
	fs, err := ioutil.ReadDir(dir)
