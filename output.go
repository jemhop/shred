package main

import (
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

	errorStyle.Print(text, "\n")
}
