package main

import (
	"flag"
	"fmt"
	"github.com/koffihuguesagossadou/bungo/internal/cmd"

)

func main() {

	if err := cmd.Do(flag.Args()); err != 0 {
		fmt.Println(err)
		return
	}
	
}
