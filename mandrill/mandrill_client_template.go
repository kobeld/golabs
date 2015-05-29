package main

import (
	"github.com/kobeld/goutils"

	"github.com/keighl/mandrill"
)

func main() {

	client := mandrill.ClientWithKey("G3liCFFFJI8pw8d9HrZnXA")

	message := &mandrill.Message{
		FromEmail: "aaron@theplant.jp",
		FromName:  "Aaron Liang",
		Subject:   "Complete your Qortex registration",
	}
	message.AddRecipient("kobeld2012@gmail.com", "Kobeld", "to")

	address := `
	The Plant<br>
	Route Kamiyamacho 6F<br>
	Kamiyama-cho 5-2<br>
	Shibuya-ku, Tokyo 150-0047<br>
	Japan<br><br>
	<a href='mailto:qortex@theplant.jp'>Qortex Team</a>
	`

	message.GlobalMergeVars = mandrill.ConvertMapToVariables(map[string]string{
		"company":           "The Plant",
		"url":               "http://qortex.com",
		"list_address_html": address,
	})

	templateContent := map[string]string{"name": "testing"}
	_, _, err := client.MessagesSendTemplate(message, "1-en-activation", templateContent)
	if err != nil {
		goutils.PrintStackAndError(err)
	}
}
