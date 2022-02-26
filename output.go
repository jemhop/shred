package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

func printFiles(files []TrashInfo) {

	/*  NAME LOCATION           DATE DELETED HASH
	LATO ~/Downloads/Lato	22/05/2022			*/

	//hashing not implemented yet
	// + i want to implement dynamic data based on term size
	//e.g skinniest = name + location, then + date, then + hash

	tableData := make([][]string, 0)

	tableData = append(tableData, []string{"NAME", "PATH", "DATE DELETED"})
	for _, file := range files {
		tableData = append(tableData, []string{file.name, file.path, file.date})
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

func printFileActions(names []string, maxLines int) {

	pterm.Bold.Println("Affected Files")
	listItems := make([]pterm.BulletListItem, 0)
	dirBullet := "🗀 "
	bulletStyle := pterm.NewStyle(pterm.Bold)

	for _, name := range names {
		stat, err := os.Stat(name)
		checkErr(err)
		if stat.IsDir() {
			listItems = append(listItems, pterm.BulletListItem{Level: 0, Text: name, Bullet: dirBullet, BulletStyle: bulletStyle})
		} else {
			listItems = append(listItems, pterm.BulletListItem{Level: 0, Text: name, Bullet: dirBullet, BulletStyle: bulletStyle})
		}

	}

	pterm.DefaultBulletList.WithItems(listItems).Render()
}
