package main

import (
	"fmt"
	"log"

	"github.com/FahadAlothman-fsd/projector-go/pkg/cli"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatalf("unable to get opts %v", err)
	}
	fmt.Printf("opts: %+v\n", opts)
}
