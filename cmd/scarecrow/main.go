/*
Main entry point for the Scarecrow chatbot application.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	scarecrow "scarecrow"
	_ "scarecrow/listeners/console"
	_ "scarecrow/listeners/slack"
	_ "scarecrow/listeners/xmpp"
)

func main() {
	// Collect command line parameters.
	debug := flag.Bool("debug", false, "Enable debug logging.")
	version := flag.Bool("version", false, "Show the version number and exit.")
	flag.Parse()

	if *version == true {
		fmt.Printf("This is Scarecrow, version %s\n", scarecrow.VERSION)
		os.Exit(0)
	}

	// Create the bot instance.
	bot := scarecrow.New()
	bot.Debug = *debug

	bot.Start()
}
