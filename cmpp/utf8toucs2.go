package main

import (
	"fmt"

	cmpputils "github.com/bigwhite/gocmpp/utils"
)

func main() {

	content, err := cmpputils.Utf8ToUcs2("MGUA")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}
