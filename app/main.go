package main

import (
	"fmt"

	toolkit "github.com/cicorias/udemny-toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)

	fmt.Println("Random String: ", s)
}
