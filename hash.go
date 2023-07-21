package helpers

import (
	"crypto/md5"
	"fmt"
)

// MD5 computes the md5 checksum of the given data.
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
