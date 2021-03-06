package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//this file hosts everything involved in interacting with the XDG trash specification, followed by gnome + kde

type TrashInfo struct {
	name   string
	path   string
	date   string
	folder bool
}

func getTrashList() []TrashInfo {
	trashInfoDir := filepath.Join(getTrashDir(), "info")

	unfiliteredFiles, err := os.ReadDir(trashInfoDir)
	checkErr(err)

	files := make([]string, 0)

	infoArr := make([]TrashInfo, len(files))

	for _, file := range unfiliteredFiles {
		if filepath.Ext(file.Name()) == ".trashinfo" {
			files = append(files, file.Name())
		}
	}

	for _, file := range files {
		infoArr = append(infoArr, getTrashInfoFile(filepath.Join(trashInfoDir, file)))
	}

	return infoArr
}

func getTrashInfoFile(path string) TrashInfo {
	infoFile, err := os.Open(path)
	checkErr(err)

	defer infoFile.Close()
	info := readTrashLines(infoFile)

	filePath := getFilePathFromInfoPath(path)

	stat, err := os.Stat(filePath)
	checkErr(err)

	info.folder = stat.IsDir()
	return info
}

func getFilePathFromInfoPath(path string) string {
	name := filepath.Base(path)
	return filepath.Join(getTrashDir(), "files", name[0:len(name)-len(".trashinfo")])
}

//this whole function is very hardcoded, but im not gonna bother supporting things not compliant with freedesktop specs
// besides maybe macOS support in future
func readTrashLines(file *os.File) TrashInfo {
	sc := bufio.NewScanner(file)
	path, date := "", ""
	lastLine := 0
	for sc.Scan() {
		lastLine++
		if lastLine == 2 {
			path = sc.Text()
		}
		if lastLine == 3 {
			date = sc.Text()
		}
	}
	path = strings.Split(path, "=")[1]
	date = strings.Split(date, "=")[1]
	name := filepath.Base(file.Name()[0 : len(file.Name())-len(".trashinfo")])
	return TrashInfo{name: name, path: path, date: date}
}

func getTrashDir() string {
	dirname, err := os.UserHomeDir()
	checkErr(err)

	return dirname + "/.local/share/Trash"
}

func trashExists(name string) bool {
	list := getTrashList()

	for _, f := range list {
		if name == f.name {
			return true
		}
	}
	return false
}

func createTrashInfo(filename string, origin string) {
	//freedesktop trash spec uses RFC3339 format
	currentTime := time.Now().Format(time.RFC3339)
	currentTime = currentTime[0 : len(currentTime)-len("+00:00")]

	filename += ".trashinfo"
	f, err := os.Create(filepath.Join(getTrashDir(), "info", filepath.Base(filename)))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	//according to the spec this filepath should be encoded according to rfc 2396, but even dolphin isnt fully compliant with this
	//it only encodes spaces and nothing else as far as i can tell, but also has no problem restoring from a non encoded string
	//this is shown as SHOULD in the spec, so im not too bothered. if it causes problems later ill figure it out properly
	//origin = url.QueryEscape(origin)

	f.WriteString("[Trash Info]\n")
	f.WriteString("Path=" + origin + "\n")
	f.WriteString("DeletionDate=" + currentTime)

}
