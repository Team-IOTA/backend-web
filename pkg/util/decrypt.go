package util

import (
    "classifie/pkg/models"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	return gcm.Open(nil, nonce, ciphertext, nil)
}

func GetDecryptObject(object models.Model_User)(models.Model_User , error){

	    var err error
        var decryptedFeild []byte
    	key := []byte("the-key-has-to-be-32-bytes-long!")
    	decryptedFeild ,err = Decrypt([]byte(object.Password ), key)
    	 object.Password = Bytes2StrRaw(decryptedFeild)
    	
    	if err != nil{
    	    return object,err
    	}
        return object,nil
}

func GetDecryptObjects(objects []models.Model_User)([]models.Model_User , error){

	    var err error
        var decryptedFeild []byte
		var decryptedObjects []models.Model_User

    	key := []byte("the-key-has-to-be-32-bytes-long!")
		for i :=0; i < len(objects);i++{
			decryptedFeild ,err = Decrypt([]byte(objects[i].Password ), key)
			objects[i].Password = Bytes2StrRaw(decryptedFeild)
			
			if err != nil{
				return objects,err
			}
			decryptedObjects = append(decryptedObjects,objects[i])
		}

        return decryptedObjects,nil
}
