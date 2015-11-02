package postmark_test

import (
	"fmt"

	"github.com/tikiatua/postmark"
)

func ExampleService() {

	// initialize the service (a default host will be used if no Host is given)
	mailings := postmark.Service{
		APIKey: "YOUR API KEY",
	}

	// compose a new message
	message := postmark.Message{
		From: "SENDER EMAIL",
		To:   []string{"RECIPIENT EMAIL"},

		Subject:  "SUBJECT",
		TextBody: "MESSAGE-BODY-AS-TEXT",
	}

	// send the message through the postmark service
	response, err := mailings.Send(&message)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", response)
}
