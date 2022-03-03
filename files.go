package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func nOverwriteRuns(file string, n int) {
	for i := 0; i < n; i++ {
		overwriteFile(file)
	}
}

//This function parses _ as if it is selecting every file in the trash directory if trashDir is set to true
func getArgsFromPath(args []string, trashDir bool) []string {
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

func overwriteFile(path string) {
	chunkSize := int64(16000000)

	stat, err := os.Stat(path)
	checkErr(err)

	//file mode doesn't really matter, new file creation flag not set
	f, err := os.OpenFile(path, os.O_RDWR, os.FileMode(0000))
	checkErr(err)

	size := stat.Size()
	defer f.Close()

	// if file > 16mb, overwrite in chunks to avoid storing excessive amounts of data in memory at once
	// if this is stupid and or pointless lmk
	offset := int64(0)
	for size > 0 {
		currentSize := Min(size, chunkSize)

		f.WriteAt(nRandomBytes(currentSize), 0)

		offset += currentSize
		size -= currentSize
	}
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
			checkErr(err)
		}

	}

	return output
}

//This function will check for duplicates and append a number to the end of the file name to avoid conflicts
func getSafeMovePath(origin string, dest string) string {
	dest, err := url.QueryUnescape(dest)
	checkErr(err)

	if fileExists(dest, false) {
		for i := 1; true; i++ {
			ext := filepath.Ext(dest)
			noExt := dest[0 : len(dest)-len(ext)]

			suffix := " (" + strconv.Itoa(i) + ")"
			if !fileExists(dest+suffix, false) && i > 0 {
				return noExt + suffix + ext
			}
		}
	}
	return dest
}
