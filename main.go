package main

import (
	"custom-lang/src"
	"flag"
	"log"
)

func main() {
	var filename *string
	filename = flag.String("main", "", "Name of the main file")
	var dryRun *bool
	dryRun = flag.Bool("dryRun", false, "Only load code, useful to see errors")

	flag.Parse()
	if *filename == "" {
		log.Fatal("You must provide a file name\n")
	}

	//var program *execution.Program
	var err error
	_, err = src.LoadModule(*filename)

	if err != nil {
		log.Fatal(err)
	}

	if !*dryRun {
		// TODO : Run program here
	}
}
