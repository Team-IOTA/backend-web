package util

import (
    "classifie/pkg/models"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
func Bytes2StrRaw(b []byte) string {
	return string(b[:])
}

func GetEncryptObject (object models.Model_User)(models.Model_User,error){
    var err error
    var encryptedFeild []byte
	key := []byte("the-key-has-to-be-32-bytes-long!")
	encryptedFeild ,err = Encrypt([]byte(object.Password ), key)
	 object.Password = Bytes2StrRaw(encryptedFeild)
	
	if err != nil{
	    return object,err
	}
    return object,nil
}