package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Data struct {
	GenerateSign []int64 `json:"generateSign"`
	VerifySign   []int64 `json:"verifySign"`
}

var result1 []int64
var result2 []int64

func main() {
	for i := 0; i < 10; i++ {

		// The GenerateKey method takes in a reader that returns random bits, and
		// the number of bits
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			panic(err)
		}

		publicKey := privateKey.PublicKey
		readData, err := os.ReadFile("test4.json")
		if err != nil {
			fmt.Println("error is", err)
		}
		msg, err := json.Marshal(readData)
		if err != nil {
			fmt.Println("error is", err)
		}
		// msg := []byte("verifiable message")

		// Before signing, we need to hash our message
		// The hash is what we actually sign
		msgHash := sha256.New()
		_, err = msgHash.Write(msg)
		if err != nil {
			panic(err)
		}
		msgHashSum := msgHash.Sum(nil)

		// In order to generate the signature, we provide a random number generator,
		// our private key, the hashing algorithm that we used, and the hash sum
		// of our message
		start1 := time.Now()
		signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
		if err != nil {
			panic(err)
		}
		result1 = append(result1, time.Since(start1).Microseconds())
		fmt.Println("time of generate sign is", result1[i])
		// To verify the signature, we provide the public key, the hashing algorithm
		// the hash sum of our message and the signature we generated previously
		// there is an optional "options" parameter which can omit for now
		start2 := time.Now()
		err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
		if err != nil {
			fmt.Println("could not verify signature: ", err)
			return
		}
		result2 = append(result2, time.Since(start2).Microseconds())
		fmt.Println("time of verify is", result2[i])
		// If we don't get any error from the `VerifyPSS` method, that means our
		// signature is valid
		fmt.Println("signature verified")

	}
	data := Data{
		GenerateSign: result1,
		VerifySign:   result2,
	}
	output, err := json.MarshalIndent(&data, "", "\t\t")
	if err != nil {
		fmt.Println("转换失败")

	}
	err = os.WriteFile("result4.json", output, 0644)
	if err != nil {
		fmt.Println("写文件失败")
	}
}
