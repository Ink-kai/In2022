package utils

import (
	"crypto/md5"
	"encoding/binary"
	"strconv"
	"time"

	"github.com/golang-module/dongle"
)

func GenerateID() uint64 {
	nowdate := time.Now().Unix()
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(nowdate, 10)))
	return binary.BigEndian.Uint64(h.Sum(nil))
}

// AES加密
func EncryptPassword(str string) []byte {
	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)        // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)   // No、Zero、PKCS5、PKCS7
	cipher.SetKey("1234567887654321") // key 长度必须是 16、24 或 32
	cipher.SetIV("1234567887654321")  // iv 长度必须是 16、24 或 32
	cypt := dongle.Encrypt.FromBytes([]byte(str)).ByAes(cipher).ToHexBytes()
	return cypt
}

//AES解密
func DecodePassword(encrypt string) []byte {
	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)                                                     // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)                                                // No、Zero、PKCS5、PKCS7
	cipher.SetKey("1234567887654321")                                              // key 长度必须是 16、24 或 32
	cipher.SetIV("1234567887654321")                                               // iv 长度必须是 16、24 或 32
	cyptde := dongle.Decrypt.FromHexBytes([]byte(encrypt)).ByAes(cipher).ToBytes() // []byte("hello world")
	return cyptde
}
