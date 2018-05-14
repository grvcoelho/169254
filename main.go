package main

import (
	"fmt"

	"github.com/grvcoelho/169254/server"
)

func debug(x interface{}) {
	fmt.Printf("%T: %#v\n", x, x)
}

func main() {
	server := server.New()

	server.Start()
}
