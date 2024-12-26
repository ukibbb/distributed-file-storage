package main

import (
	"crypto/rand"
	"encoding/hex"
)

func generateID() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}
