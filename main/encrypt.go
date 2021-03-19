// This file contains encrypt function

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

// encryptFile encrypts the given and put it into ./encrypt folder
func encryptFile(key string, inFileHandler, outFileHandler *os.File) error {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := inFileHandler.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			_, err = outFileHandler.Write(buf[:n])
			if err != nil {
				return err
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}

	_, err = outFileHandler.Write(iv)
	return err
}
