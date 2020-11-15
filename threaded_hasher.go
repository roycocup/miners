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
	var result = make(chan string)

	for found == false {
		var res = go hashIt(secret, nonce, targetStr)

		counter = counter + 1
		nonce = nonce + 1

		if counter >= maxIterations {
			fmt.Println("Max iterations reached!")
			break
		}

	}

	fmt.Printf("Hash: %s\n", result)
	fmt.Printf("Iterations: %d\n", nonce)

}

func hashIt(secret string, nonce int, targetStr string) (int, string, string){
	h := sha256.New()
	strNonce := strconv.Itoa(nonce)
	var noncedSecret = []byte(secret + strNonce)
	h.Write(noncedSecret)
	var hexResult = h.Sum(nil)

	result := hex.EncodeToString(hexResult)

	if wasFound(targetStr, result){
		return 1, result, strNonce
	}
	return 0, "", ""
}

func wasFound(targetStr string, result string) bool {
	if strings.Compare(result[:len(targetStr)], targetStr[:len(targetStr)]) == 0 {
		return true
	}
	return false
}
