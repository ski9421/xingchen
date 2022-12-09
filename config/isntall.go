package config

import (
	"fmt"
	"os"
)

func CretaeLock() {
	// 创建文件,返回两个值,一个是创建的文集爱你,二是错误信息
	f, err := os.Create("./config/install.lock")
	if err != nil { //如果有错误 打印错误 返回
		fmt.Println("err=", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
}

func IsNotLock() (bool, error) {
	_, err := os.Stat("./config/install.lock")
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
