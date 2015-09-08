package main

import (
	"fmt"
	"net"
)

func main() {
	//-------------------------
	service := "xmpp-server"
	protocol := "tcp"
	name := "google.com"

	cname, addresses, err := net.LookupSRV(service, protocol, name)

	if err != nil {
		panic(err)
	}

	fmt.Printf("cname : %s \n", cname)

	for i := 0; i < len(addresses); i++ {

		fmt.Printf("addrs[%d].Target : %s \n", i, addresses[i].Target)
		fmt.Printf("addrs[%d].Port : %d \n", i, addresses[i].Port)
		fmt.Printf("addrs[%d].Priority : %d \n", i, addresses[i].Priority)
		fmt.Printf("addrs[%d].Weight : %d \n", i, addresses[i].Weight)
	}

}
