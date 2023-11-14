package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	testType, exists := os.LookupEnv("STOREGH_TEST_TYPE")
	if !exists {
		log.Fatalf("missing required \"TYPE\" field to determine test type")
	}

	log.Printf("Running %q test type\n", testType)

	switch testType {
	case "LOAD_TEST_TOKEN":
		loadTypeTokenTestCheck()

	case "LOAD_TEST_APP":
		loadTypeAppTestCheck()

	case "UPDATE_TEST_TOKEN":
		updateTypeTokenTestCheck()

	case "UPDATE_TEST_APP":
		updateTypeAppTestCheck()

	default:
		log.Fatalf("received unsupported test type")
	}
}

func loadTypeTokenTestCheck() {
	// Token Load Test 1
	envVar1 := "STOREGH_REPOSITORY_ACTIONS_TEST_TOKEN_LOAD_KEY_1_BAD_SYMBOLS________"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar1)
	}
	if secret != "I am secret one!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		log.Printf("the value of %q matched the expected value! Yay!\n", envVar1)
	}

	// Token Load Test 2
	envVar2 := "STOREGH_REPOSITORY_ACTIONS_TEST_TOKEN_LOAD_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar2)
	}
	if secret != "I am secret two!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		log.Printf("the value of %q matched the expected value! Yay!\n", envVar2)
	}

	// Token Load Test 3
	envVar3 := "STOREGH_REPOSITORY_ACTIONS_TEST_TOKEN_LOAD_KEY_3"
	secret, found = os.LookupEnv(envVar3)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar3)
	}
	if secret != "1234" {
		log.Fatalf("received %q value but did not match expected value\n", envVar3)
	} else {
		log.Printf("the value of %q matched the expected value! Yay!\n", envVar3)
	}

	// Token Load Test 4
	envVar4 := "STOREGH_REPOSITORY_ACTIONS_TEST_TOKEN_LOAD_KEY_4"
	secret, found = os.LookupEnv(envVar4)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar4)
	}
	if secret != "{\"0\":\"my\",\"1\":\"map\",\"2\":\"value\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar4)
	} else {
		var val any
		if err := json.Unmarshal([]byte(secret), &val); err != nil {
			log.Fatalf("received %q value but it could not be unmarshalled into valid JSON", envVar4)
		}

		log.Printf("the value of %q matched the expected value! Yay!\n", envVar4)
	}

	// Token Load Test 5
	envVar5 := "STOREGH_REPOSITORY_ACTIONS_TEST_TOKEN_LOAD_KEY_5"
	secret, found = os.LookupEnv(envVar5)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar5)
	}
	if secret != "{\"number\":1,\"true\":false}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar5)
	} else {
		var val any
		if err := json.Unmarshal([]byte(secret), &val); err != nil {
			log.Fatalf("received %q value but it could not be unmarshalled into valid JSON", envVar5)
		}

		log.Printf("the value of %q matched the expected value! Yay!\n", envVar5)
	}
}

func loadTypeAppTestCheck() {
	// App Load Tets 1
	envVar5 := "STOREGH_REPOSITORY_ACTIONS_TEST_APP_LOAD_KEY_1"
	secret, found := os.LookupEnv(envVar5)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar5)
	}
	if secret != "12345" {
		log.Fatalf("received %q value but did not match expected value\n", envVar5)
	} else {
		log.Printf("the value of %q matched the expected value! Yay!\n", envVar5)
	}
}

func updateTypeTokenTestCheck() {
	// Token Update Test 1
	envVar1 := "STOREGH_REPOSITORY_ACTIONS_TEST_TOKEN_UPDATE_KEY_1"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar1)
	}
	if secret != "value-1-updated" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		log.Printf("the value of %q matched the expected value! Yay!\n", envVar1)
	}

	// Token Update Test 2
	envVar2 := "STOREGH_REPOSITORY_ACTIONS_TEST_TOKEN_UPDATE_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar2)
	}
	if secret != "{\"0\":\"value-2-updated\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		log.Printf("the value of %q matched the expected value! Yay!\n", envVar2)
	}
}

func updateTypeAppTestCheck() {
	// App Update Test 1
	envVar3 := "STOREGH_REPOSITORY_ACTIONS_TEST_APP_UPDATE_KEY_1"
	secret, found := os.LookupEnv(envVar3)
	if !found || secret == "" {
		log.Fatalf("Failed to find %q value\n", envVar3)
	}
	if secret != "value-1-updated" {
		log.Fatalf("received %q value but did not match expected value\n", envVar3)
	} else {
		log.Printf("the value of %q matched the expected value! Yay!\n", envVar3)
	}
}
