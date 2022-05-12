package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"io"
	"math/big"
	"sort"
	"strings"
)

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var presetKey = bytes.NewBufferString("0CoJUm6Qyw8W8jud")
var iv = bytes.NewBufferString("0102030405060708")

const publicKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----"
const eapiKey = "e82ckenh8dichen8"

func WeapiCrypto(obj interface{}) map[string]string {
	t, _ := json.Marshal(obj)
	text := string(t)
	text = "{\"phone\":\"17671880651\",\"countrycode\":\"86\",\"password\":\"9486ba8669ef0af6e4b04615a78f2431\",\"rememberLogin\":\"true\",\"csrf_token\":\"\"}"
	randomBytes := make([]byte, 16)
	for i, v := range make([]byte, 16) {
		randomBytes[i] = byte(base62[v%62])
	}

	var secretKey = bytes.NewBuffer(randomBytes)

	temp := base64.StdEncoding.EncodeToString(AesEncrypt(bytes.NewBufferString(text), "cbc", presetKey, iv))
	params := base64.StdEncoding.EncodeToString(AesEncrypt(bytes.NewBufferString(temp), "cbc", secretKey, iv))

	encSecKey := hex.EncodeToString(RsaEncrypt(secretKey.Bytes(), publicKey))

	res := map[string]string{
		"params":    params,
		"encSecKey": encSecKey,
	}
	return res
}

func EapiCryptp(url string, obj interface{}) map[string]string {
	t, _ := json.Marshal(obj)
	text := string(t)
	message := "nobody" + url + "use" + text + "md5forencrypt"
	m := md5.New()
	m.Write([]byte(message))
	digest := hex.EncodeToString(m.Sum(nil))

	data := url + "-36cd479b6b5-" + text + "-36cd479b6b5-" + digest

	params := hex.EncodeToString(AesEncrypt(bytes.NewBufferString(data), "ecb", bytes.NewBufferString(eapiKey), nil))

	return map[string]string{
		"params": strings.ToTitle(params),
	}
}

//--------------------aes---------------------------------------------
func AesEncrypt(origData *bytes.Buffer, mode string, key *bytes.Buffer, iv *bytes.Buffer) []byte {

	var encrypted []byte
	switch mode {
	case "cbc":
		encrypted = aesEncryptCBC(origData.Bytes(), key.Bytes(), iv.Bytes())
	case "ecb":
		encrypted = aesEncryptECB(origData.Bytes(), key.Bytes())
	case "cfb":
		encrypted = aesEncryptCFB(origData.Bytes(), key.Bytes())
	}

	return encrypted

}

func aesEncryptCBC(origData []byte, key []byte, iv []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                 // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)   // 补全码
	blockMode := cipher.NewCBCEncrypter(block, iv) // 加密模式
	encrypted = make([]byte, len(origData))        // 创建数组
	blockMode.CryptBlocks(encrypted, origData)     // 加密
	return encrypted
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func aesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipherT, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipherT.BlockSize(); bs <= len(origData); bs, be = bs+cipherT.BlockSize(), be+cipherT.BlockSize() {
		cipherT.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func aesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}

//--------------------rsa----------------------------------
func RsaEncrypt(planText []byte, key string) []byte {
	sort.SliceStable(planText, func(i, j int) bool {
		return true
	})

	pt := bytes.NewBuffer(make([]byte, 128-len(planText)))
	pt.Write(planText)

	buf := bytes.NewBufferString(key)
	//pem解码
	block, _ := pem.Decode(buf.Bytes())
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey2 := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	c := new(big.Int).SetBytes(pt.Bytes())
	encryptedBytes := c.Exp(c, big.NewInt(int64(publicKey2.E)), publicKey2.N).Bytes()

	return encryptedBytes

}
