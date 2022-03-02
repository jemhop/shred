package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/atomicgo/cursor"
	"github.com/pterm/pterm"
)

func printFiles(files []TrashInfo) {

	/*  NAME LOCATION           DATE DELETED HASH
	LATO ~/Downloads/Lato	22/05/2022			*/

	//hashing not implemented yet
	// + i want to implement dynamic data based on term size
	//e.g skinniest = name + location, then + date, then + hash

	tableData := make([][]string, 0)

	tableData = append(tableData, []string{"TYPE", "NAME", "PATH", "DATE DELETED"})
	for _, file := range files {
		prefix := ""
		if file.folder {
			prefix = "dir"
		} else {
			prefix = "file"
		}

		tableData = append(tableData, []string{prefix, file.name, file.path, file.date})
	}
	pterm.DefaultTable.WithBoxed(true).WithHasHeader().WithHeaderStyle(pterm.NewStyle(pterm.FgRed, pterm.Bold)).WithData(tableData).Render()
}

func printError(text string) {
	errorStyle := pterm.NewStyle(pterm.FgRed, pterm.Bold, pterm.Underscore)

	errorStyle.Println(text)
}

func printYesNoPrompt(text string, yesDefault bool) bool {
	checkbox := ""
	if yesDefault {
		checkbox = " [Y/n]"
	} else if !yesDefault {
		checkbox = " [y/N]"
	} else {
		log.Fatal("Invalid default selection provided to yes/no prompt, provide 0 or 1")
	}

	promptStyle := pterm.NewStyle(pterm.FgWhite, pterm.Italic)
	promptStyle.Print(text)
	pterm.Bold.Println(checkbox)

	input := ""
	fmt.Scanln(&input)

	input = strings.ToLower(input)
	input = strings.TrimSpace(input)

	if input == "" {
		if yesDefault {
			return true
		} else {
			return false
		}
	}

	if input == "y" || input == "yes" {
		return true
	} else {
		return false
	}
}

func printFileActions(names []string, maxLines int, doIndent bool, showDirs bool) {
	pterm.Bold.Println("Affected Files")
	listItems := make([]pterm.BulletListItem, 0)
	excludedDirs := 0

	for _, name := range names {
		stat, err := os.Stat(name)
		checkErr(err)
		indent := 0
		if doIndent {
			indent = strings.Count(name[0:len(name)-len(filepath.Base(name))], "/")
		}

		if !stat.IsDir() {
			listItems = append(listItems, pterm.BulletListItem{Level: indent, Text: name, Bullet: "🗍", BulletStyle: pterm.NewStyle(pterm.Bold)})
		} else if showDirs {
			listItems = append(listItems, pterm.BulletListItem{Level: indent, Text: name, Bullet: "🗀", BulletStyle: pterm.NewStyle(pterm.Bold)})
		} else {
			excludedDirs++
		}

		if len(listItems) == maxLines {
			break
		}
	}

	pterm.DefaultBulletList.WithItems(listItems).Render()

	remaining := len(names[len(listItems):])
	if remaining-excludedDirs > 0 {
		cursor.Up(1)
		pterm.Italic.Println(" ... and " + strconv.Itoa(remaining-excludedDirs) + " other files (folders not counted) \n")
	}

}
