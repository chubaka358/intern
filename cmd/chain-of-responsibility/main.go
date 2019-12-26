package main

import (
	"fmt"

	. "github.com/chubaka358/intern/pkg/chain-of-responsibility"
)

func main() {
	handlers := GetHandler{&DeleteHandler{&PostHandler{&PutHandler{}}}}
	fmt.Println("%q", handlers.SendRequest("GT", "1"))
}
