package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//creating sub-commands.
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//creating flags for each sub-commands.
	firstFlag := subCommand1.Bool("processing", false, "command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Golang", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)
	}
}
