package main

import (
	"donkey-admin/lib/libutils"
	"fmt"
)

func main() {
	pwd := "demo111"
	pwdHash := libutils.EncryptWord(pwd)
	fmt.Println(pwdHash)
}
