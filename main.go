package main

import (
	_ "github.com/joho/godotenv/autoload" // Keep as first!

	"fmt"
	"github.com/nuriofernandez/keys-in-door-experiment/httpKeys"
	"time"
)

func main() {
	fmt.Println("Starting...")
	go httpKeys.Start()
	for {
		fmt.Println("Checking if keys are in the door...")
		there, _ := AreKeysThere()
		httpKeys.KeysInDoor = there
		time.Sleep(1 * time.Second)
	}
}
