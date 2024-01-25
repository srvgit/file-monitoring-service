package main

import (
	internal "file-monitoring-service/internal/processor"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	path := os.Getenv("CONFIG_PATH")

	fmt.Println("Loading config from ", path)
	processor, err := internal.NewProcessor()

	processor.ProcessEvents()

	if err != nil {
		log.Fatal(err)
	}

}
