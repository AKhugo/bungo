package main

import (
	"github.com/koffihuguesagossadou/bungo/pkg/fi"
	"github.com/koffihuguesagossadou/bungo/pkg/format"
	"fmt"
	"flag"
	"os"
)

func main() {
	
	// user input file
	inputFile := flag.String("i", "", "input file")
	outputFile := flag.String("o", "", "output file")

	flag.Parse()

	if *inputFile == "" {
		fmt.Println("input file is required")
		flag.Usage() // display the help message
		return
	}

	// check if file exist
	data, err := fi.GetFileData(*inputFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	// encode file in base64
	encoded, err := format.EncodeToBase64(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	// write encoded file
	if *outputFile != "" {
		err := os.WriteFile(*outputFile, []byte(encoded), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		
	}

	fmt.Println(encoded)

}
