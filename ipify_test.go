package main

import (
	"log"
	"testing"
)

func TestGetMyIP(t *testing.T) {
	ip, err := getMyIP(-1)

	if err != nil {
		t.Error(err)
	}

	if len(ip) == 0 {
		t.Error("Empty ip")
	}

	log.Println("IP " + ip)
}
