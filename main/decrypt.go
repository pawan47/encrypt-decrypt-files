// This file contains decrypt function

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)

// decryptFile decrypt the given file and put it into ./decrypt folder
func decryptFile(key string, inFileHandler, outFileHandler *os.File) error {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	fi, err := inFileHandler.Stat()
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	msgLen := fi.Size() - int64(len(iv))
	_, err = inFileHandler.ReadAt(iv, msgLen)
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := inFileHandler.Read(buf)
		if n > 0 {
			if n > int(msgLen) {
				n = int(msgLen)
			}
			msgLen -= int64(n)
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
	return nil
}
