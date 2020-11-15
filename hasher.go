package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var nonce = 1
	var maxIterations = 10000000
	var targetStr = "000000"
	var secret = "once this is found out there is not more"

	var found = false
	var counter = 0
	var result = ""

	for found == false {
		h := sha256.New()
		var noncedSecret = []byte(secret + strconv.Itoa(nonce))
		h.Write(noncedSecret)
		var hexResult = h.Sum(nil)

		result = hex.EncodeToString(hexResult)

		if strings.Compare(result[:len(targetStr)], targetStr[:len(targetStr)]) == 0 {
			found = true
		}

		counter = counter + 1
		nonce = nonce + 1

		if counter > maxIterations {
			fmt.Println("Max iterations reached!")
			break
		}

	}

	fmt.Printf("Hash: %s\n", result)
	fmt.Printf("Iterations: %d\n", nonce)

}
