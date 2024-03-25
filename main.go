package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload" // Keep as first!
	"github.com/nuriofernandez/keys-in-door-experiment/httpKeys"
)

func main() {
	fmt.Println("Starting...")
	httpKeys.Start()
}
