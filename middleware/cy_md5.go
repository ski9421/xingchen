package middleware

import (
	"crypto/md5"
	"fmt"
)

// 生成Md5

func CreateMd5(str string) string {
	sum := md5.Sum([]byte(str))
	s := fmt.Sprintf("%x", sum)
	return s
}
