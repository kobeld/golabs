package main

import (
	"fmt"
	"strings"
)

func main() {
	old := "The Plant"
	// newStr := strings.Replace(old, " ", "-", -1)
	// println(newStr)

	// groupName := strings.Replace(organization.Name, " ", "-", -1)
	groupName := strings.Replace(old, " ", "-", -1)

	cmd := fmt.Sprintf("srg_create %v %v %v %v %v", groupName, "kobeld2.local", groupName, groupName, groupName)
	println(cmd)

	mail := "aaron@theplant.jp"
	sa := strings.Split(mail, "@")
	sa[1] = strings.Replace(sa[1], ".", "-", -1)
	word := strings.Join(sa, "-")

	println(word)
}
