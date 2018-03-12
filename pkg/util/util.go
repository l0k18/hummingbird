// Package util - miscellaneous utility functions for Hummingbird PoW
package util

import (
	"crypto/rand"
	"encoding/binary"
	"log"
)

// GetRandUint64 - gets a cryptographically secure pseudorandom byte slice and converts it to uint32
func GetRandUint64() uint64 {
	r := make([]byte, 8)
	n, err := rand.Read(r)
	if n != 8 {
		log.Fatal("For some reason unable to get 8 random bytes")
	}
	if err != nil {
		log.Fatal(err)
	}
	return BytesToUint64(r)
}

// Uint64ToBytes - returns a byte slice from uint64 - required because Murmur3 takes bytes as input but returns uint32
func Uint64ToBytes(input uint64) []byte {
	p := make([]byte, 8)
	binary.LittleEndian.PutUint64(p, input)
	return p
}

// BytesToUint64 - converts 4 byte slice to uint32
func BytesToUint64(bytes []byte) uint64 {
	if len(bytes) != 8 {
		log.Fatal("Byte slice is not 8 bytes long")
	}
	return binary.LittleEndian.Uint64(bytes)
}

// Head - takes a uint646 and returns the first 16 bits
func Head(input uint64) uint32 {
	return uint32((input << 32) >> 32)
}

// Tail - takes a uint64 and returns the second 16 bits
func Tail(input uint64) uint32 {
	return uint32(input >> 32)
}
