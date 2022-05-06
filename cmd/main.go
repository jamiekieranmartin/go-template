package main

import (
	"flag"
	"fmt"
)

const cliVersion = "0.0.1"

const helpMessage = `
go-template v%s

`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("go-template v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}
}
