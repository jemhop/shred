package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

func main() {
	args := os.Args

	//im just about 100% sure this is an illegal way of handling cmdline arguments but i hate programming this kind of really basic io,
	//which ends up being a pain in the ass because turning arbitrary text input into a series of actions where each action is contextually
	//decided by all the other bits of arbitrary text input is not fun in my personal opinon
	if len(args) == 1 {
		invalid()
	}

	//TODO: support * to select all files in directory/restore all files, support more than one file/folder per arg (scan further in args list)
	for i, arg := range args {
		if i != 0 {
			if arg[0] == '-' {
				if arg == "-h" {
					help()
					break
				} else if arg == "-l" {
					list()
					break
				} else if arg == "-r" {
					valid := getAllPathArgs(args[i+1:], true)
					recoverFiles(valid)
					break
				} else if arg == "-d" {
					valid := getAllPathArgs(args[i+1:], false)
					deleteFiles(valid)
					break
				} else if arg == "-s" {
					valid := getAllPathArgs(args[i+1:], false)
					shredFiles(valid)
					break
				} else if arg == "-e" {
					valid := getAllPathArgs(args[i+1:], true)
					eraseFiles(valid)
				} else {
					invalid()
				}
			}
		}
	}

}

func recoverFiles(names []string) {
	for _, name := range names {
		spinner, _ := pterm.DefaultSpinner.Start("Checking " + name + " is in trash")

		if !trashExists(name) {
			spinner.Fail("File name not found")
			return
		}

		spinner.UpdateText("Getting move path")

		infoPath := filepath.Join(getTrashDir(), "info", name+".trashinfo")
		trashInfo := getTrashInfo(infoPath)
		origin := filepath.Join(getTrashDir(), "files", trashInfo.name)
		destination := getSafePath(origin, trashInfo.path)

		spinner.UpdateText("Moving file to " + destination)

		os.Rename(origin, destination)
		os.Remove(infoPath)

		spinner.Success(name + " has been moved to " + destination)
	}
}

func deleteFiles(names []string) {
	for _, name := range names {
		spinner, _ := pterm.DefaultSpinner.Start("Checking " + name + " exists")

		workingDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		if !fileExists(name, false) {
			spinner.Fail("File name not found")
			return
		}

		spinner.UpdateText("Getting safe move path")

		uncheckedDest := filepath.Join(getTrashDir(), "files", filepath.Base(name))

		dest := getSafePath(name, uncheckedDest)

		spinner.UpdateText("Moving file to trash")

		os.Rename(name, dest)
		createTrashInfo(filepath.Base(dest), filepath.Join(workingDir, name))

		spinner.Success(name + " has been moved to trash")
	}

}

func shredFiles(names []string) {

	files := filesFromDirs(names, true)

	printFileActions(files, 5)

	pterm.NewStyle(pterm.FgRed, pterm.Bold).Println("This deletion cannot be undone by any process, including professional drive recovery")
	if printYesNoPrompt("Are you sure you want to delete these files permanently and irretrievably?", false) {
		for _, file := range files {

			spinner, _ := pterm.DefaultSpinner.Start("Shredding " + file)

			stat, err := os.Stat(file)
			if err != nil {
				spinner.Fail(file + " doesn't exist, skipping")
				continue
			}

			if stat.Size() == 0 || stat.IsDir() {
				spinner.Warning(file + "is 0b or directory, can safely delete")
			} else {
				shred(file)
				spinner.Success(file + "succesfully shredded")
			}
			os.Remove(file)
		}
	}
}

func eraseFiles(names []string) {
	for _, name := range names {
		spinner, _ := pterm.DefaultSpinner.Start("Checking " + name + " is in trash")

		if !trashExists(name) {
			spinner.Fail("File name not found")
			return
		}

		spinner.UpdateText("Getting move path")

		infoPath := filepath.Join(getTrashDir(), "info", name+".trashinfo")
		trashInfo := getTrashInfo(infoPath)
		origin := filepath.Join(getTrashDir(), "files", trashInfo.name)

		spinner.UpdateText("Deleting file from trash")

		os.Remove(origin)
		os.Remove(infoPath)

		spinner.Success(name + " has been permanently deleted")

	}
}

func list() {
	files := getTrashList()
	printFiles(files)
}

func invalid() {
	printError("Use -h to see valid arguments")
}

func help() {
	pterm.DefaultTable.WithBoxed(true).WithHasHeader().WithHeaderStyle(pterm.NewStyle(pterm.FgRed, pterm.Bold)).WithData(pterm.TableData{
		{"Command", "Usage"},
		{"l (list)", "Lists all deleted files"},
		{"r (restore)", "Recover a file by name (_ will restore all)"},
		{"d (delete)", "Deletes a file by moving to the trash"},
		{"s (shred)", "Overwrites a file 5x and deletes permanently "},
		{"e (erase)", "Deletes a file in the trash permanently"},
	}).Render()
}
