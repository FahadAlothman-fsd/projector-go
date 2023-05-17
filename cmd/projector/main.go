package main

import (
	"encoding/json"
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

	// fmt.Printf("opts: %+v\n", opts)
	// fmt.Printf("config: %+v\n", config)

	proj := projector.NewProjector(config)
	switch config.Operation {
	case projector.Print:
		{
			if len(config.Args) == 0 {
				data := proj.GetValueAll()
				jsonString, err := json.Marshal(data)
				if err != nil {
					log.Fatalf("unable to marshal data %v\n", err)
				}
				fmt.Printf("%v\n", string(jsonString))

			} else if value, ok := proj.GetValue(config.Args[0]); ok {
				fmt.Printf("%v\n", value)
			}
		}
	case projector.Add:
		{
			proj.SetValue(config.Args[0], config.Args[1])
			proj.Save()
		}
	case projector.Remove:
		{
			proj.RemoveValue(config.Args[0])
			proj.Save()
		}
	}

}
