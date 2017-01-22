# msg card

print card package

# example

```
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
```

```
$ go run example/example.go 
┌──────────┐
│foo       │
├──────────┤
│bar       │
│bar       │
│baz       │
│aaaaaaaaaa│
│aaaaaa    │
└──────────┘
```
