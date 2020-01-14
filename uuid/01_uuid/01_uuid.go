package main

import (
	"log"

	"github.com/gofrs/uuid"
)

func main() {
	// Create a Version 4 UUID.
	u1, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u1)

	// Convert to byte array
	a1 := u1.Bytes()
	log.Printf("convert UUID to bytes %v", a1)

	// Reverse byte array to UUID
	u2, err := uuid.FromBytes(a1)
	if err != nil {
		log.Fatalf("failed to parse UUID from bytes %q: %v", a1, err)
	}

	log.Printf("successfully parsed UUID %v", u2)

	// Parse a UUID from a string.
	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	u3, err := uuid.FromString(s)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", s, err)
	}
	log.Printf("successfully parsed UUID %v", u3)
}
