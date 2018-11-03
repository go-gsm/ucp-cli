package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/go-gsm/ucp"
)

func main() {
	opt := &ucp.Options{
		Addr:       fmt.Sprintf("%s:%s", os.Getenv("SMSC_HOST"), os.Getenv("SMSC_PORT")),
		User:       os.Getenv("SMSC_USER"),
		Password:   os.Getenv("SMSC_PASSWORD"),
		AccessCode: os.Getenv("SMSC_ACCESSCODE"),
	}

	client := ucp.New(opt)
	if err := client.Connect(); err != nil {
		fmt.Println("Cant connect")
		os.Exit(1)
	}
	defer client.Close()

	reader, _ := readline.New(">>> ")
	defer reader.Close()

	for {
		fmt.Print(">>> ")
		lines, _ := reader.Readline()
		fields := strings.Fields(lines)

		if len(fields) == 1 {
			// exit CLI
			if fields[0] == "exit" {
				return
			}
			// display help message
			if fields[0] == "help" {
				fmt.Println("\n\tSend a 'message' to 'receiver' with a 'sender' mask\n\t>>> sender receiver message\n\n\tExit the cli\n\t>>> exit\n")
			}
		}

		// sender receiver message...
		if len(fields) >= 3 {
			sender := fields[0]
			receiver := fields[1]
			message := strings.Join(fields[2:], " ")
			ids, err := client.Send(sender, receiver, message)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v\n", ids)
			}
		}
	}
}
