package main

import (
	"fmt"

	card "github.com/ieee0824/msg_card"
)

func main() {
	msg := card.MessageBox{
		"foo",
		"bar\nbar\nbaz\naaaaaaaaaaaaaaaa",
	}

	fmt.Println(msg.String(10))
}
