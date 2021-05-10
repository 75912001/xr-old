package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//GenMd5 生成md5
func GenMd5(data []byte) (md5sum string) {
	md5hash := md5.New()
	md5hash.Write(data)
	cipherStr := md5hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MD5File(pathFile string) (md5sum string, err error) {
	f, err := os.Open(pathFile)
	if err != nil {
		return
	}
	defer f.Close()

	md5hash := md5.New()
	_, err = io.Copy(md5hash, f)
	if err != nil {
		return
	}
	md5sum = hex.EncodeToString(md5hash.Sum(nil))
	return
}
