package crypt

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"math/rand"
	"strconv"
	"time"

	"github.com/jeffotoni/organizando.seu.projeto.go/projeto2/internal/fmts"
)

var SALT = "#!.*&393939393%4..48484!%38383"

//Sha1 ..
func Sha1(key string) string {
	data := []byte(fmts.Concat(key, SALT))
	return (fmt.Sprintf("%x", sha1.Sum(data)))
}

func Crc32() string {
	crc32q := crc32.MakeTable(0xD5828281)
	return fmt.Sprintf("%08x", crc32.Checksum([]byte(RandomNew()), crc32q))
}

func RandomNew() string {
	rand.Seed(time.Now().UnixNano()) //UnixNano
	return Md5(strconv.Itoa(rand.Intn(1000000) + rand.Intn(100000)))
}

func Md5(text string) string {
	h := md5.New()
	io.WriteString(h, text)
	return (fmt.Sprintf("%x", h.Sum(nil)))
}
