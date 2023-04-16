package main

import "flag"

var flags struct {
	showGenerated    bool
	showNotGenerated bool
}

func parse() {
	flag.BoolVar(&flags.showGenerated, "generated", false, "Log about functions that was generated.")
	flag.BoolVar(&flags.showNotGenerated, "not-generated", true, "Log about functions that was NOT generated.")
}
