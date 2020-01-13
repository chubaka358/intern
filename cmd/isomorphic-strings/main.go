package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/isomorphic-strings"
)

func main() {
	isomorph := isomorphic_strings.NewIsomorph()
	fmt.Println(isomorph.IsIsomorphic("ab", "aa"))
}
