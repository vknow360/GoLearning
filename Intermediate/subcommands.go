package main

import "flag"

func main() {
	subCommand := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subCommand.Bool("processing", false, "Command processing status")
	secondFlag := subCommand2.Int("level", 0, "Processing level")

	flag.Parse()

	if len(flag.Args()) < 1 {
		println("This program requires additional commands")
		return
	}
	switch flag.Args()[0] {
	case "firstSub":
		subCommand.Parse(flag.Args()[1:])
		if *firstFlag {
			println("First subcommand processing is enabled.")
		}
	case "secondSub":
		subCommand2.Parse(flag.Args()[1:])
		println("Second subcommand processing level:", *secondFlag)
	}

}
