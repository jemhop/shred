package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func shred(file string) {
	for i := 0; i < 5; i++ {
		randomOverwriteFile(file)
	}
	os.Remove(file)
}

//This function parses _ as if it is selecting every file in the trash directory if trashDir is set to true
func getAllPathArgs(args []string, trashDir bool) []string {
	valid := make([]string, 0)

	for _, arg := range args {
		if trashDir {
			if filepath.Base(arg) == "_" && !trashExists(filepath.Base(arg)) {
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

func randomOverwriteFile(path string) {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	//file mode doesn't really matter, new file will never be created
	f, err := os.OpenFile(path, os.O_RDWR, os.FileMode(0707))
	if err != nil {
		log.Fatal(err)
	}

	size := stat.Size()
	defer f.Close()

	f.WriteAt(nRandomBytes(size), 0)

}

//takes a list of files, and directories and returns a list of all files
func filesFromDirs(names []string, includeDirs bool) []string {
	output := make([]string, 0)

	for _, name := range names {
		stat, err := os.Stat(name)
		checkErr(err)

		if !stat.IsDir() {
			output = append(output, name)
		} else {
			err := filepath.Walk(name,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						fmt.Println(err)
						return err
					}
					if !info.IsDir() || includeDirs {
						output = append(output, path)
					}
					return nil
				})
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	return output
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
