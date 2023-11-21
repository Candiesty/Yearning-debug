package dmessage

import (
	"log"
)

// 是否显示调试信息
var showDM = true

func PrintV(a ...any) {
	if showDM {
		for _, each := range a {
			log.Printf("%+v\n", each)
		}
	}
}
