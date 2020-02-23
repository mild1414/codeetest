package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					*files = append(&files, f.Name())
				}
			}

			if f.IsDir() {
				Path := dir + "/" + f.Name()
				*Path = append(*Path, Path, dir+"/"+f.Name())
				FindFileFromExtension(extension, path, files, Path)
			}
		}
	}
}

func main() {
	drives := getDrives := range drives {
		FindFileFromExtension([]string{".txt"}, rangedrives, &files, &Path)
	}
	file, err := os.Create("output.txt")
}
