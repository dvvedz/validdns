package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	validdnsservers "github.com/dvvedz/validdns"
)

func main() {
	var inputFileFlag string
	flag.StringVar(&inputFileFlag, "", "", "a string")
	flag.Parse()

	if len(inputFileFlag) != 0 {
		f, fileErr := os.Open(inputFileFlag)
		if fileErr != nil {
			log.Fatalf("error while opening file: %s, err: %s", inputFileFlag, fileErr)
		}
		defer f.Close()
	}

	result := string(validdnsservers.GetPublicDnsServers())
	fmt.Println(result)

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// line := scanner.Text()
	// fmt.Println(line)
	// }
}
