package utils

import (
	"crypto/md5"
	"fmt"
)

const Md5Key = "u8r7XR1z"

func Md5(data string) string {
	// 进行md5加密，因为Sum函数接受的是字节数组，因此需要注意类型转换
	srcCode := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", srcCode)
}

func Md5Encryption(text string) string {
	return Md5(Md5Key + text)
}
