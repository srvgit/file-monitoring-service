package main

import (
	internal "file-monitoring-service/internal/processor"

	log "github.com/sirupsen/logrus"
)

func main() {

	processor, err := internal.NewProcessor()

	processor.ListenFolderEvents()

	if err != nil {
		log.Fatal(err)
	}

}
