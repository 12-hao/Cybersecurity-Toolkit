package utils

import "os" //判断是否存在该文件

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path) //用来判断这个文件是否存在
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
