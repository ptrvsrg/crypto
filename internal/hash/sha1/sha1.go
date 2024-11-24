package sha1

import (
	"encoding/binary"
	"math/bits"
)

// Initialize values for SHA-1
var h0, h1, h2, h3, h4 uint32

// Functions for working with data blocks
func padding(input []byte) []byte {
	length := len(input)
	paddingLength := (56 - (length+1)%64) % 64
	padded := append(input, 0x80) // Add bit 1 (0x80)

	// Pad with zeros
	for i := 0; i < paddingLength; i++ {
		padded = append(padded, 0x00)
	}

	// Add data length
	bitLength := uint64(length * 8)
	for i := 0; i < 8; i++ {
		padded = append(padded, byte(bitLength>>(56-i*8)))
	}

	return padded
}

func processBlock(block []byte) {
	// Get a 16-word array
	w := make([]uint32, 80)
	for i := 0; i < 16; i++ {
		w[i] = binary.BigEndian.Uint32(block[i*4 : (i+1)*4])
	}

	// Expand to 80 words
	for i := 16; i < 80; i++ {
		w[i] = bits.RotateLeft32(w[i-3]^w[i-8]^w[i-14]^w[i-16], 1)
	}

	// Initializing Variables
	a, b, c, d, e := h0, h1, h2, h3, h4

	// Main SHA-1 loop
	for i := 0; i < 80; i++ {
		var f, k uint32
		if i < 20 {
			f = (b & c) | (^b & d)
			k = 0x5A827999
		} else if i < 40 {
			f = b ^ c ^ d
			k = 0x6ED9EBA1
		} else if i < 60 {
			f = (b & c) | (b & d) | (c & d)
			k = 0x8F1BBCDC
		} else {
			f = b ^ c ^ d
			k = 0xCA62C1D6
		}

		temp := bits.RotateLeft32(a, 5) + f + e + k + w[i]
		e = d
		d = c
		c = bits.RotateLeft32(b, 30)
		b = a
		a = temp
	}

	// Update values
	h0 += a
	h1 += b
	h2 += c
	h3 += d
	h4 += e
}

func Sum(input []byte) []byte {
	// Initialize hash state
	h0 = 0x67452301
	h1 = 0xEFCDAB89
	h2 = 0x98BADCFE
	h3 = 0x10325476
	h4 = 0xC3D2E1F0

	// Add padding
	paddedInput := padding(input)

	// Process 512-bit blocks
	for i := 0; i < len(paddedInput); i += 64 {
		processBlock(paddedInput[i : i+64])
	}

	// Compose the result
	hash := make([]byte, 20)
	binary.BigEndian.PutUint32(hash[0:4], h0)
	binary.BigEndian.PutUint32(hash[4:8], h1)
	binary.BigEndian.PutUint32(hash[8:12], h2)
	binary.BigEndian.PutUint32(hash[12:16], h3)
	binary.BigEndian.PutUint32(hash[16:20], h4)

	return hash
}
