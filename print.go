package main

import (
	"github.com/robmonte/private-printer-restricted/privateprinterrestricted"
	"github.com/robmonte/private-printer/privateprinter"
)

func Print() {
	privateprinter.PrivatePrint("dependabot test with access")
	privateprinterrestricted.PrivatePrint("dependabot test without access")
}
