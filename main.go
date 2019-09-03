package main

import (
	"changeip/telnet"
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
)

var spin = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

func main() {
	initializeConfig()
	result, err := askForIPPrefix()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Show spinner
	spin.Start()

	// Get current ip
	initialIP, err := getMyIP(-1)

	if err != nil {
		log.Fatalln(err)
	}

	// Create telnet client
	tcClient, err := telnet.New()

	if err != nil {
		log.Fatalln(err)
		return
	}

	// Start process
	log.Println("Initial IP " + initialIP)
	newIP := changeIPTo(tcClient, initialIP, result)
	log.Println("NEW IP " + newIP)
}
