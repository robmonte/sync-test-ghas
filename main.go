package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	envVar1 := "sinkerAWSSM_Test_Key_1_Bad_Symbols________"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar1)
	}

	if secret != "I am secret one!" {
		log.Fatalf("Received %q value but did not match expected value\n", envVar1)
	} else {
		fmt.Printf("The value of %q matched the expected value! Yay!\n", envVar1)
	}

	envVar2 := "sinkerGHAS-Test-Secret-2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar2)
	}

	if secret != "I am secret one!" {
		log.Fatalf("Received %q value but did not match expected value\n", envVar2)
	} else {
		fmt.Printf("The value of %q matched the expected value! Yay!\n", envVar2)
	}
}
