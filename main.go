package main

import (
	"fmt"
	"log"
	"os"
	//"github.com/robmonte/private-printer/privateprinter"
)

func main() {
	// privateprinter.PrivatePrint("PRINTING!!!!!")

	testType, exists := os.LookupEnv("TYPE")
	if !exists {
		log.Fatalf("missing required \"TYPE\" field to determine test type")
	}

	fmt.Printf("Running %q test type\n", testType)

	switch testType {
	case "LOAD_TEST_TOKEN":
		loadTypeTokenTestCheck()

	case "LOAD_TEST_APP":
		loadTypeAppTestCheck()

	case "UPDATE_TEST_TOKEN":
		updateTypeTokenTestCheck()

	default:
		log.Fatalf("received unsupported test type")
	}
}

func loadTypeTokenTestCheck() {
	envVar1 := "STOREGHAS_TEST_LOAD_KEY_1_BAD_SYMBOLS________"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar1)
	}

	if secret != "I am secret one!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar1)
	}

	envVar2 := "STOREGHAS_TEST_LOAD_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar2)
	}

	if secret != "I am secret two!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar2)
	}

	envVar3 := "STOREGHAS_TEST_LOAD_KEY_3"
	secret, found = os.LookupEnv(envVar3)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar3)
	}

	if secret != "1234" {
		log.Fatalf("received %q value but did not match expected value\n", envVar3)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar3)
	}

	envVar4 := "STOREGHAS_TEST_LOAD_KEY_4"
	secret, found = os.LookupEnv(envVar4)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar4)
	}

	if secret != "{\"0\":\"my\",\"1\":\"map\",\"2\":\"value\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar4)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar4)
	}
}

func loadTypeAppTestCheck() {
	envVar5 := "STOREGHAS_TEST_LOAD_KEY_5"
	secret, found := os.LookupEnv(envVar5)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar5)
	}

	if secret != "12345" {
		log.Fatalf("received %q value but did not match expected value\n", envVar5)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar5)
	}
}

func updateTypeTokenTestCheck() {
	envVar1 := "STOREGHAS_TEST_UPDATE_KEY_1"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar1)
	}

	if secret != "StoreGHAS-Test-Update-Value-1-Changed" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar1)
	}

	envVar2 := "STOREGHAS_TEST_UPDATE_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar2)
	}

	if secret != "{\"0\":\"StoreGHAS-Test-Update-Value-2-Changed\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar2)
	}
}

func updateTypeAppTestCheck() {
	envVar3 := "STOREGHAS_TEST_UPDATE_KEY_3"
	secret, found := os.LookupEnv(envVar3)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar3)
	}

	if secret != "StoreGHAS-Test-Update-Value-3-Changed" {
		log.Fatalf("received %q value but did not match expected value\n", envVar3)
	} else {
		fmt.Printf("the value of %q matched the expected value! Yay!\n", envVar3)
	}
}
