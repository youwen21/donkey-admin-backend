package main

import (
	"fmt"
	"gofly/lib/libutils"
)

func main() {
	pwd := "admin"
	pwdHash := libutils.EncryptWord(pwd)
	fmt.Println(pwdHash)
}
