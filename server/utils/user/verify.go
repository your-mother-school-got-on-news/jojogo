package user

import (
	"crypto/sha256"
	"reflect"
)

func Verify(str1 string, str2 string) bool {
	h := sha256.New()
	h.Write([]byte(str1))

	if reflect.DeepEqual(h.Sum(nil), []byte(str2[:])) {
		return true
	}
	return false
}
