package main

import (
	"fmt"

	"github.com/robmonte/private-printer-restricted/privateprinterrestricted"
	"github.com/robmonte/private-printer/privateprinter"
)

const INCREMENT_ME int = 1

func Print() {
	privateprinter.PrivatePrint(fmt.Sprintf("dependabot test with access %d", INCREMENT_ME))
	privateprinterrestricted.PrivatePrint(fmt.Sprintf("dependabot test no access %d", INCREMENT_ME))
}
