package service

import (
	"crypto/md5"
	"fmt"
)
 func Md5Password(str string)(md5string string){
	md5string = fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return md5string
}

