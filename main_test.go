package main

import (
	"bytes"
	"testing"
)

func TestMACParser(t *testing.T) {
	mac := "68:37:e9:ae:9d:85"
	macBytes := []byte{0x68, 0x37, 0xE9, 0xAE, 0x9D, 0x85}

	if !bytes.Equal(macBytes, MACParser(mac)) {
		t.Errorf("Failed MAC conversion")
	}
}
