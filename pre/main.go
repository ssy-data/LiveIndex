package pre

import (
	"encoding/json"
	"fmt"
	"os"
	"pre/curve"
	"pre/recrypt"
	"time"
)

type Data struct {
	GenerateKey   []int64 `json:"generateKey"`
	EncryptFile   []int64 `json:"encryptFile"`
	ReEncryptFile []int64 `json:"reEncryptFile"`
	DecryptFile   []int64 `json:"decryptFile"`
	Redecrypt     []int64 `json:"redecrypt"`
}

func TestRecrypt() {
	var result1 []int64
	var result2 []int64
	var result3 []int64
	var result4 []int64
	var result5 []int64
	for i := 0; i < 10; i++ {
		// Alice Generate Alice key-pair
		aPriKey, aPubKey, _ := curve.GenerateKeys()
		fmt.Println("Generate Alice key-pair")
		// Bob Generate Bob key-pair
		bPriKey, bPubKey, _ := curve.GenerateKeys()
		fmt.Println("Generate Bob key-pair")
		start1 := time.Now()
		rk, pubX, err := recrypt.ReKeyGen(aPriKey, bPubKey)
		if err != nil {
			fmt.Println(err)
		}
		result1 = append(result1, time.Since(start1).Microseconds())
		// result1[i] = time.Since(start1).Microseconds()
		fmt.Println("generate re-encryption key time is", time.Since(start1).Microseconds())
		fmt.Println("generates re-encryption key rk:", rk)
		start2 := time.Now()
		fileCapsule, err := recrypt.EncryptFile("dataset/hash/hash4.txt", "dataset/hash/hash4_encrypt.txt", aPubKey)
		if err != nil {
			fmt.Println("File Encrypt Error:", err)
		}
		result2 = append(result2, time.Since(start2).Microseconds())
		fmt.Println("EncryptFile time is", time.Since(start2).Microseconds())
		fmt.Println("EncryptFile over")
		fileData, err := json.Marshal(fileCapsule)
		if err != nil {
			fmt.Println("ReEncryption Error:", err)
		}
		fmt.Println("fileCapsule is", string(fileData))
		start3 := time.Now()
		fileNewCapsule, err := recrypt.ReEncryption(rk, fileCapsule)
		if err != nil {
			fmt.Println("ReEncryption Error:", err)
		}
		result3 = append(result3, time.Since(start3).Microseconds())
		fmt.Println("ReEncryptFile time is", time.Since(start3).Microseconds())
		fmt.Println("ReEncryptFile over")
		start4 := time.Now()
		time5 := recrypt.DecryptFile("dataset/hash/hash4_encrypt.txt", "dataset/hash/hash4_decrypt.txt", bPriKey, fileNewCapsule, pubX)
		if err != nil {
			fmt.Println("Decrypt Error:", err)
		}
		result4 = append(result4, time.Since(start4).Microseconds())
		fmt.Println("decryptFile time is", time.Since(start4).Microseconds())
		fmt.Println("DecryptFile over")
		result5 = append(result5, time5)
		fmt.Println("redecrypt time is", time5)

	}
	data := Data{
		GenerateKey:   result1,
		EncryptFile:   result2,
		ReEncryptFile: result3,
		DecryptFile:   result4,
		Redecrypt:     result5,
	}
	output, err := json.MarshalIndent(&data, "", "\t\t")
	if err != nil {
		fmt.Println("转换失败")

	}
	err = os.WriteFile("result8.json", output, 0644)
	if err != nil {
		fmt.Println("写文件失败")
	}
}
