// This file contains main function or entry point of package

// ask if encrypt or decrypt
// then will ask key - if entered then uses default key
// then ask for the file to encrypt or decrypt
// saves a new file with suffix _decypted or _encrypted acc to operation

// WIll be using
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		key           string
		filePath      string
		outFilepath   string
		err           error
		opToDoStr     string
		opIDInt       int
		fileHandler   *os.File
		outFileHandle *os.File

		defaultKey = "1234567891123456"
	)

	defer func(err error) {
		if err != nil {
			_ = fmt.Errorf("rolling back and deleting any file if created\n")
			_ = fmt.Errorf("error occured %w", err)
		}
	}(err)

	fmt.Printf("Enter 1 to decrypt and 2 to encrypt: ")
	opToDoStr = scanStr()
	opIDInt, err = strconv.Atoi(opToDoStr)
	if err != nil {
		return
	}

	fmt.Printf("Enter your key string: ")
	key = padPassKey(scanStr())
	if key == "" {
		key = defaultKey
	}

	fmt.Printf("Enter filepath to operate: ")
	filePath = scanStr()
	if filePath == "" {
		err = errors.New("no file path entered exiting code")
		return
	}
	fileHandler, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer fileHandler.Close()

	fmt.Printf("Enter output filepath: ")
	outFilepath = scanStr()
	if outFilepath == "" {
		err = errors.New("no file path entered exiting code")
		return
	}
	outFileHandle, err = os.Open(outFilepath)
	if err != nil {
		return
	}
	defer outFileHandle.Close()

	switch opIDInt {
	case 1:
		err = decryptFile(key, fileHandler, outFileHandle)
	case 2:
		err = encryptFile(key, fileHandler, outFileHandle)
	default:
		err = errors.New("wrong opID entered")
	}
	return
}
