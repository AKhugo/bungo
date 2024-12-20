package main

import (
	"flag"
	"fmt"

	"github.com/koffihuguesagossadou/bungo/pkg/command"
)

func main() {

	if err := command.Do(flag.Args()); err != 0 {
		fmt.Println(err)
		return
	}
	
}
