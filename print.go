package main

import (
	"github.com/robmonte/private-printer/privateprinter"
)

const INCREMENT_ME int = 1

func Print() {
	privateprinter.PrivatePrint("dependabot test!")
}
