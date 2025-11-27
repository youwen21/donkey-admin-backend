package main

import (
	"fmt"
	"gofly/lib/libutils"
)

func main() {
	pwd := "demo111"
	pwdHash := libutils.EncryptWord(pwd)
	fmt.Println(pwdHash)
}
