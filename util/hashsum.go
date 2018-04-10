package util

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
)

// VerifyHashsum func
func VerifyHashsum(file1, file2 []byte) (bool, error) {
	sh1, sh2 := sha512.New(), sha512.New()

	if _, err := io.Copy(sh1, bytes.NewReader(file1)); err != nil {
		log.Println(err)
		return false, err
	}

	if _, err := io.Copy(sh2, bytes.NewReader(file2)); err != nil {
		log.Println(err)
		return false, err
	}

	return fmt.Sprintf("%x", sh1.Sum(nil)) == fmt.Sprintf("%x", sh2.Sum(nil)), nil
}

// HashAndVerify hashes given bytearray and verify with given hash
func HashAndVerify(file []byte, hash string) (bool, error) {
	sh := sha512.New()

	if _, err := io.Copy(sh, bytes.NewReader(file)); err != nil {
		log.Println(err)
		return false, err
	}

	hash1 := fmt.Sprintf("%x", sh.Sum(nil))

	return hash == hash1, nil
}
