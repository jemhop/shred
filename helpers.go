package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func fileExists(path string, print bool) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else if print {
			log.Fatal(err)
		}
	}
	return true
}

//This function will check for duplicates and append a number to the start of the file name to avoid conflicts
func getSafePath(origin string, dest string) string {
	suffix := ""
	i := 0

	dest, err := url.QueryUnescape(dest)

	if err != nil {
		log.Fatal(err)
	}

	for {
		extension := filepath.Ext(dest)
		noExtension := dest[0 : len(dest)-len(extension)]
		if fileExists(dest+suffix, false) {
			conv := strconv.Itoa(i + 1)
			suffix = " (" + conv + ")"
		} else {
			return noExtension + suffix + extension
		}
		i++
	}
}

//This function parses * as if it is selecting every file in the trash directory
func getAllPathArgs(args []string, trashDir bool) []string {
	valid := make([]string, 0)

	for _, arg := range args {
		if trashDir {
			fmt.Println(filepath.Base(arg))
			if filepath.Base(arg) == "_" {
				trash := getTrashList()
				for _, t := range trash {
					valid = append(valid, t.name)
				}
			} else {
				valid = append(valid, arg)
			}
		} else {
			valid = append(valid, arg)
		}
	}

	return valid
}
