package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"

	// "encoding/base64"
	b64 "encoding/base64"
	// "encoding/hex"

	"io"

	"github.com/Gearbox-protocol/sdk-go/log"
)

func getByte32Key(key []byte) []byte {
	h := sha256.New()
	h.Write([]byte(key))
	return h.Sum(nil)
}

func byteToBase64(msg []byte) string {
	sEnc := b64.StdEncoding.EncodeToString(msg)
	return sEnc
}
func base64ToByte(sEnc string) []byte {
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	return sDec
}

func Encrypt(message, pass string) string {
	text := []byte(message)
	key := getByte32Key([]byte(pass))
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	log.CheckFatal(err)

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	log.CheckFatal(err)

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	_, err = io.ReadFull(rand.Reader, nonce)
	log.CheckFatal(err)

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	encryptedMsgInByte := gcm.Seal(nonce, nonce, text, nil)
	encodedStr := byteToBase64(encryptedMsgInByte)
	return encodedStr
}

func Decrypt(encryptedMsg string, pass []byte) string {

	key := getByte32Key(pass)
	ciphertext := base64ToByte(encryptedMsg)
	// if our program was unable to read the file
	// print out the reason why it can't
	c, err := aes.NewCipher(key)
	log.CheckFatal(err)

	gcm, err := cipher.NewGCM(c)
	log.CheckFatal(err)

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Fatal("cypher len is more than blockSize(nonceSize)")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	log.CheckFatal(err)
	return string(plaintext)
}
