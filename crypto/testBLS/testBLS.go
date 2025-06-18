package testBLS

import (
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Nik-U/pbc"
)

var BLS_sys *BLS

type BLS struct {
	bls_pairing *pbc.Pairing
	bls_g       *pbc.Element
}
type BLS_byte struct {
	Params string
	G      []byte
}

var blsfilePath = "./bls.gob"

type Data struct {
	GenerateSign []int64 `json:"generateSign"`
	VerifySign   []int64 `json:"verifySign"`
}

var result1 []int64
var result2 []int64

func main() {
	for i := 0; i < 10; i++ {
		Save_bls_sys_into_lib()
		BLS_start()
		readData, err := os.ReadFile("test4.json")
		if err != nil {
			fmt.Println("error is", err)
		}
		msg, err := json.Marshal(readData)
		if err != nil {
			fmt.Println("error is", err)
		}
		privKey, pubKey := Generate_bls_keypair()
		start1 := time.Now()
		signature := Bls_signature(msg, privKey)
		result1 = append(result1, time.Since(start1).Microseconds())
		fmt.Println("time of generate sign is", result1[i])
		sibyte := SetSIGIntoByte(signature)
		sign := SetPubKeyFromByte(sibyte)
		start2 := time.Now()
		Bls_verify(msg, pubKey, sign)
		result2 = append(result2, time.Since(start2).Microseconds())
		fmt.Println("time of verify is", result2[i])

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

func SetPubKeyFromByte(pubkey []byte) *pbc.Element {
	return BLS_sys.bls_pairing.NewG2().SetBytes(pubkey)
}
func SetSIGIntoByte(sig *pbc.Element) []byte {
	return sig.Bytes()
}

func Save_bls_sys_into_lib() {
	params := pbc.GenerateA(160, 512)
	para_byte := params.String()
	g_byte := params.NewPairing().NewG2().Rand().Bytes()
	bls_byte := BLS_byte{para_byte, g_byte}
	blsSaveToFilebyte(bls_byte)
}
func BLS_start() {
	blsbyte_ := blsLoadFromFile()
	pairing, _ := pbc.NewPairingFromString(blsbyte_.Params)
	g := pairing.NewG2().SetBytes(blsbyte_.G)
	BLS_sys = &BLS{pairing, g}

}
func blsSaveToFilebyte(bls BLS_byte) {
	pa := &bls
	file, _ := os.OpenFile(blsfilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	enc.Encode(pa)
}
func blsLoadFromFile() *BLS_byte {
	file, _ := os.Open(blsfilePath)
	defer file.Close()
	var pa BLS_byte
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	return &pa
}
func Generate_bls_keypair() (*pbc.Element, *pbc.Element) {
	privKey := BLS_sys.bls_pairing.NewZr().Rand()
	//g := _pairing.NewG2().Rand()
	pubKey := BLS_sys.bls_pairing.NewG2().PowZn(BLS_sys.bls_g, privKey)
	return privKey, pubKey
}
func Bls_signature(message []byte, privkey *pbc.Element) *pbc.Element {
	h := BLS_sys.bls_pairing.NewG1().SetFromStringHash(string(message), sha256.New())
	signature := BLS_sys.bls_pairing.NewG2().PowZn(h, privkey)
	return signature
}
func Bls_verify(message []byte, pubkey *pbc.Element, signature *pbc.Element) bool {
	h := BLS_sys.bls_pairing.NewG1().SetFromStringHash(string(message), sha256.New())
	temp1 := BLS_sys.bls_pairing.NewGT().Pair(h, pubkey)
	temp2 := BLS_sys.bls_pairing.NewGT().Pair(signature, BLS_sys.bls_g)
	if !temp1.Equals(temp2) {
		fmt.Println("*BUG* Signature check failed *BUG*")
		return false
	} else {
		fmt.Println("*Signature check  success*")
		return true
	}
}
