package main

import (
	"fmt"

	"github.com/keighl/mandrill"
)

func main() {

	client := mandrill.ClientWithKey("G3liCFFFJI8pw8d9HrZnXA")

	message := &mandrill.Message{}
	message.AddRecipient("kobeld2012@gmail.com", "Kobeld", "to")
	message.FromEmail = "aaron@theplant.jp"
	message.FromName = "Aaron Liang"
	message.Subject = "You won the prize!"
	message.HTML = "<h1>You won!!</h1>"
	message.Text = "You won!!"

	responses, apiError, err := client.MessagesSend(message)
	if err != nil {
		println(err.Error())
	}

	if apiError != nil {
		fmt.Printf("\n %+v \n", apiError)
	}

	for _, resp := range responses {
		fmt.Printf("\n %+v \n", resp)
	}

}
