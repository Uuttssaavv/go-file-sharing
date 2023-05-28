package utils

import (
	"crypto/rand"
	"unsafe"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(length int) string{
	b := make([]byte, length)
    rand.Read(b)
    for i := 0; i < length; i++ {
        b[i] = alphabet[b[i] % byte(len(alphabet))]
    }
    return *(*string)(unsafe.Pointer(&b))
}