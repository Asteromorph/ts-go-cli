package main

import (
	"fmt"
	"log"

	"github.com/Asteromorph/aoc/pkg/projector"
)

func main() {
    opts, err := projector.GetOpts()
    if err != nil {
        log.Fatalf("unable to get options %v", err)
    }

    config, err := projector.NewConfig(opts) 
    if err != nil {
        log.Fatalf("unable to get configs %v", err)
    }
    fmt.Printf("configs: %+v", config)
}
