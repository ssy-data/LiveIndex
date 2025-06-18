package SM2

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/tjfoc/gmsm/sm2"
)

type Data struct {
	GenerateKey []int64 `json:"generateKey"`
	Encrypt     []int64 `json:"encrypt"`
	Decrypt     []int64 `json:"Decrypt"`
}

func TestSignatureSM2() {
	for i := 0; i < 10; i++ {
		//生成私钥
		privateKey, e := sm2.GenerateKey(rand.Reader)
		if e != nil {
			fmt.Println("sm2 encrypt faild")
		}
		//从私钥中获取公钥
		pubkey := &privateKey.PublicKey
		readData, err := os.ReadFile("test4.json")
		if err != nil {
			fmt.Println("error is", err)
		}
		msg, err := json.Marshal(readData)
		if err != nil {
			fmt.Println("error is", err)
		}
		//用公钥加密msg
		bytes, err := pubkey.EncryptAsn1(msg, rand.Reader)

		if err != nil {
			fmt.Println("使用私钥加密失败！")
		}

		fmt.Println("the encrypt msg  =  ", hex.EncodeToString(bytes))
		//用私钥解密msg
		decrypt, err := privateKey.DecryptAsn1(bytes)

		if err != nil {

			fmt.Println("使用私钥解密失败！")
		}
		fmt.Println("the msg  = ", string(decrypt))

	}

}
