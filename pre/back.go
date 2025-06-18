package pre

// func main() {
// 	// Alice Generate Alice key-pair
// 	aPriKey, aPubKey, _ := curve.GenerateKeys()
// 	fmt.Println("Generate Alice key-pair")
// 	// Bob Generate Bob key-pair
// 	bPriKey, bPubKey, _ := curve.GenerateKeys()
// 	fmt.Println("Generate Bob key-pair")
// 	// plain text
// 	// m := "Hello, Proxy Re-Encryption"
// 	// fmt.Println("origin message:", m)
// 	// Alice encrypts to get cipherText and capsule
// 	// cipherText, capsule, err := recrypt.Encrypt(m, aPubKey)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// capsuleAsBytes, err := recrypt.EncodeCapsule(*capsule)
// 	// if err != nil {
// 	// 	fmt.Println("encode error:", err)
// 	// }
// 	// capsuleTest, err := recrypt.DecodeCapsule(capsuleAsBytes)
// 	// if err != nil {
// 	// 	fmt.Println("decode error:", err)
// 	// }
// 	// fmt.Println("capsule before encode:", capsule)
// 	// fmt.Println("capsule after decode:", capsuleTest)
// 	// fmt.Println("ciphereText:", cipherText)
// 	// Test recreate aes key
// 	// keyBytes, err := recrypt.RecreateAESKeyByMyPriKey(capsule, aPriKey)
// 	// if err != nil {
// 	// 	fmt.Println("Recreate key error:", err)
// 	// }
// 	// fmt.Println("recreate key:", hex.EncodeToString(keyBytes))
// 	// Alice generates re-encryption key
// 	start := time.Now()
// 	rk, pubX, err := recrypt.ReKeyGen(aPriKey, bPubKey)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("generate re-encryption key time is", time.Since(start).Microseconds())
// 	fmt.Println("generates re-encryption key rk:", rk)
// 	// Server executes re-encrypt
// 	// newCapsule, err := recrypt.ReEncryption(rk, capsule)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// // Bob decrypts the cipherText
// 	// plainText, err := recrypt.Decrypt(bPriKey, newCapsule, pubX, cipherText)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// plainTextByMyPri, err := recrypt.DecryptOnMyPriKey(aPriKey, capsule, cipherText)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println("PlainText by my own private key:", string(plainTextByMyPri))
// 	// // get plainText
// 	// fmt.Println("plainText:", string(plainText))
// 	start1 := time.Now()
// 	fileCapsule, err := recrypt.EncryptFile("test1.txt", "test1_encrypt.txt", aPubKey)
// 	if err != nil {
// 		fmt.Println("File Encrypt Error:", err)
// 	}
// 	fmt.Println("EncryptFile time is", time.Since(start1).Microseconds())
// 	fmt.Println("EncryptFile over")
// 	start2 := time.Now()
// 	fileNewCapsule, err := recrypt.ReEncryption(rk, fileCapsule)
// 	if err != nil {
// 		fmt.Println("ReEncryption Error:", err)
// 	}
// 	fmt.Println("ReEncryptFile time is", time.Since(start2).Microseconds())
// 	fmt.Println("ReEncryptFile over")
// 	start3 := time.Now()
// 	err = recrypt.DecryptFile("test1_encrypt.txt", "test1_decrypt.txt", bPriKey, fileNewCapsule, pubX)
// 	if err != nil {
// 		fmt.Println("Decrypt Error:", err)
// 	}
// 	fmt.Println("ReEncryptFile time is", time.Since(start3).Microseconds())
// 	fmt.Println("DecryptFile over")
// }
