package gost341264

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
)

var sBox = [8][16]byte{
	{0x4, 0xA, 0x9, 0x2, 0xD, 0x8, 0x0, 0xE, 0x6, 0xB, 0x1, 0xC, 0x7, 0xF, 0x5, 0x3},
	{0xE, 0xB, 0x4, 0xC, 0x6, 0xD, 0xF, 0xA, 0x2, 0x3, 0x8, 0x1, 0x0, 0x7, 0x5, 0x9},
	{0x5, 0x8, 0x1, 0xD, 0xA, 0x3, 0x4, 0x2, 0xE, 0xF, 0xC, 0x7, 0x6, 0x0, 0x9, 0xB},
	{0x7, 0xD, 0xA, 0x1, 0x0, 0x8, 0x9, 0xF, 0xE, 0x4, 0x6, 0xC, 0xB, 0x2, 0x5, 0x3},
	{0x6, 0xC, 0x7, 0x1, 0x5, 0xF, 0xD, 0x8, 0x4, 0xA, 0x9, 0xE, 0x0, 0x3, 0xB, 0x2},
	{0x4, 0xB, 0xA, 0x0, 0x7, 0x2, 0x1, 0xD, 0x3, 0x6, 0x8, 0x5, 0x9, 0xC, 0xF, 0xE},
	{0xD, 0xB, 0x4, 0x1, 0x3, 0xF, 0x5, 0x9, 0x0, 0xA, 0xE, 0x7, 0x6, 0x8, 0x2, 0xC},
	{0x1, 0xF, 0xD, 0x0, 0x5, 0x7, 0xA, 0x4, 0x9, 0x2, 0x3, 0xE, 0x6, 0xB, 0x8, 0xC},
}

func EncryptStream(in io.Reader, out io.Writer, key []uint32) error {
	buffer := make([]byte, 8)
	for {
		n, err := in.Read(buffer)
		if err != nil && err != io.EOF {
			return errors.Wrap(err, "failed to read input")
		}

		if n == 0 {
			break
		}

		if n < 8 {
			// KCS#7 padding
			padding := byte(8 - n)
			for i := n; i < 8; i++ {
				buffer[i] = padding
			}
		}

		block := binary.BigEndian.Uint64(buffer)
		encryptedBlock := encryptBlock(key, block)

		if err := writeBlock(out, encryptedBlock); err != nil {
			return errors.Wrap(err, "failed to write encrypted block")
		}

		if n < 8 {
			break
		}
	}

	return nil
}

func DecryptStream(in io.Reader, out io.Writer, key []uint32) error {
	buffer := make([]byte, 8)
	blockBytes := make([]byte, 8)

	for i := 0; ; i++ {
		n, err := in.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return errors.Wrap(err, "failed to read input")
		}

		if n == 0 {
			break
		}

		if n < 8 {
			return errors.New("input is not aligned to block size")
		}

		if i != 0 {
			if _, err := out.Write(blockBytes); err != nil {
				return errors.Wrap(err, "failed to write decrypted block")
			}
		}

		block := binary.BigEndian.Uint64(buffer)
		decryptedBlock := decryptBlock(key, block)

		binary.BigEndian.PutUint64(blockBytes, decryptedBlock)

	}

	// Delete padding
	if blockBytes != nil {
		padding := blockBytes[len(blockBytes)-1]
		if int(padding) > 8 || padding == 0 {
			return nil
		}

		// Check if padding is correct
		for i := len(blockBytes) - int(padding); i < len(blockBytes); i++ {
			if blockBytes[i] != padding {
				return nil
			}
			blockBytes[i] = 0
		}

		if _, err := out.Write(blockBytes[:len(blockBytes)-int(padding)]); err != nil {
			return errors.Wrap(err, "failed to write last block without padding")
		}
	}

	return nil
}

func encryptBlock(key []uint32, block uint64) uint64 {
	left := uint32(block >> 32)
	right := uint32(block)

	for i := 0; i < 32; i++ {
		k := key[i%8]
		temp := linearTransformation(sBoxSubstitution(modAdd(left, k)))
		left, right = right^temp, left
	}

	return (uint64(right) << 32) | uint64(left)
}

func decryptBlock(key []uint32, block uint64) uint64 {
	left := uint32(block >> 32)
	right := uint32(block)

	for i := 31; i >= 0; i-- {
		k := key[i%8]
		temp := linearTransformation(sBoxSubstitution(modAdd(left, k)))
		left, right = right^temp, left
	}

	return (uint64(right) << 32) | uint64(left)
}

func writeBlock(writer io.Writer, block uint64) error {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, block)
	_, err := writer.Write(data)
	return err
}

func modAdd(a, b uint32) uint32 {
	return (a + b) & 0xFFFFFFFF
}

func sBoxSubstitution(value uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 8; i++ {
		result |= uint32(sBox[i][(value>>(4*i))&0xF]) << (4 * i)
	}
	return result
}

func linearTransformation(value uint32) uint32 {
	return (value << 11) | (value >> (32 - 11))
}

func GenerateKey() []uint32 {
	return []uint32{
		0x01020304, 0x05060708, 0x09101112, 0x13141516,
		0x17181920, 0x21222324, 0x25262728, 0x29303132,
	}
}
