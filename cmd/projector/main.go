package main

import (
	"fmt"
	"log"

	"github.com/FahadAlothman-fsd/projector-go/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("unable to get opts %v", err)
	}
	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v\n", err)
	}

	fmt.Printf("opts: %+v\n", opts)
	fmt.Printf("config: %+v\n", config)
}
