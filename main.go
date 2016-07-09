package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	printArt()
	wait()
	runGit()
}

func printArt() {
	width, height := TerminalWidth()
	fmt.Println(BestAscii(width, height))
	fmt.Println("\n", RandomQuote())
}

func wait() {
	fmt.Println("\n\n***** Press Enter to proceed *****")
	var ignored string
	fmt.Scanln(&ignored)
}

func runGit() {
	args := os.Args[1:]

	actualCommand := "git " + strings.Join(args, " ")
	fmt.Println("Running: \n\n >", actualCommand)

	output, err := exec.Command("git", args...).CombinedOutput()

	if err != nil {
		fmt.Println(string(output), err)
	} else {
		fmt.Println(string(output))
	}
}
