package main

import (
	"log"
	"os"
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

//Returns n bytes from /dev/urandom
func nRandomBytes(num int64) []byte {
	//If my maths is right, this function is limited to reading about 9.2 petabytes from /dev/urandom due to the max length of an int64
	//I'm sure this could be overcome with an int128 type (2x int64) but 9.2 petabytes is hopefully more than anyone will ever need to overwrite
	file, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, num)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	if bytesRead == 0 {
		log.Fatal("No bytes read from /dev/urandom")
	}

	return byteSlice
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
