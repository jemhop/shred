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

func randomOverwriteFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	len := stat.Size()

	defer f.Close()

	f.WriteAt(nRandomBytes(len), 0)
}

//Returns n bytes from /dev/urandom
func nRandomBytes(num int64) []byte {
	//If my maths is right, this function is limited to reading about 9.2 petabytes from /dev/urandom due to the max length of an int64
	//I'm sure this could be overcome with an int128 type (2x int64) but 9.2 petabytes is hopefully more than anyone will ever need to overwrite
	file, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read up to len(b) bytes from the File
	// Zero bytes written means end of file
	// End of file returns error type io.EOF
	byteSlice := make([]byte, num)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
	return byteSlice
}
