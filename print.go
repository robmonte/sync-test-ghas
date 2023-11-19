package main

import (
	"github.com/robmonte/private-printer-restricted/privateprinterrestricted"
	"github.com/robmonte/private-printer/privateprinter"
)

func Print() {
	privateprinter.PrivatePrint("dependabot test with access 1")
	privateprinterrestricted.PrivatePrint("dependabot test without access 1")
}
