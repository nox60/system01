package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"reflect"
	"system01/constants"
)

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func GetEncryptedPasswd(userName string, originalPassword string) string {
	return MD5(userName + originalPassword)
}

func MD5(str string) string {
	h := md5.New()
	str = str + constants.SALTFORMD5
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
