package tools

import (
	"fmt"
)

//CheckError check the error
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
