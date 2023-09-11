package utils

import (
	"github.com/jinzhu/copier"
	"log"
)

// CopyProperties dto -> po
func CopyProperties[T any](source any) (target any) {
	if error := copier.Copy(&target, source); error != nil {
		//触发错误
		log.Panicln("CopyProperties Error Source {} Target {}")
	}
	//返回拷贝后的对象
	return target
}
