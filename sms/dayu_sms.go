package main

import (
	"fmt"

	"github.com/ltt1987/alidayu"
)

func main() {

	alidayu.AppKey = "23326361"
	alidayu.AppSecret = "685a0349223f85fcbcfb37f65f3301e6"

	success, resp := alidayu.SendSMS("18507528868", "注册验证", "SMS_5160091", `{"code":"1234","product":"Neuro"}`)
	fmt.Println("Success:", success)
	fmt.Println(resp)
}
