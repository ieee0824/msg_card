package main

import (
	"fmt"

	card "github.com/ieee0824/msg_card"
)

func main() {
	msg := card.MessageBox{
		card.String("foo"),
		card.String("特に新卒研修ではお世話になりました！今でも`エンジニア研修/事後フォロー`のページを定期的に見返しています。Facebookの投稿を今後も楽しみにしています。"),
	}

	fmt.Println(msg.String(10))
}
