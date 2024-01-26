package internal

import (
	"encoding/json"
	"file-monitoring-service/internal/config"
	"file-monitoring-service/internal/util/io"
	"fmt"
	"os"
	"path/filepath"

	"time"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

var writer chan FileEvent

var batch map[string]int64

type FileEvent struct {
	FileName string
	Size     int64
}
type Processor struct {
	config *config.Config
}

type OutPut struct {
}

type EventProcessor interface {
	NewProcessor() (*Processor, error)
	ListenFolderEvents() error
	CaptureEvents() error
}

func NewProcessor() (*Processor, error) {

	path := os.Getenv("CONFIG_PATH")
	log.Info("Loading config from ", path)
	config, err := config.LoadConfig(path)

	processor := &Processor{
		config: config,
	}
	if err != nil {
		return nil, err
	}
	return processor, nil
}

func (p *Processor) ListenFolderEvents() error {

	paths, err := io.GetSubDirectories(p.config.SrcDirectory)
	paths = append(paths, p.config.SrcDirectory)
	if err != nil {
		fmt.Println(err)
		return err
	}
	p.watch(paths)
	return nil
}

func (p *Processor) watch(paths []string) {
	if len(paths) < 1 {
		exit("must specify at least one path to watch")
	}
	// w, err := fsnotify.NewWatcher()
	w, err := fsnotify.NewBufferedWatcher(uint(p.config.MaxGoRoutines))
	workers := make(chan struct{}, p.config.MaxGoRoutines)
	writer = make(chan FileEvent)

	if err != nil {
		exit("creating a new watcher: %s", err)
	}
	defer w.Close()
	go p.CaptureEvents(w, workers)
	go p.aggregateForBatch()

	for _, p := range paths {
		err = w.Add(p)
		if err != nil {
			exit("%q: %s", p, err)
		}
	}

	printTime("ready; press ^C to exit")
	<-make(chan struct{})
}

func printTime(s string, args ...interface{}) {
	fmt.Printf(time.Now().Format("15:04:05.0000")+" "+s+"\n", args...)
}

func exit(format string, a ...interface{}) {
	log.Errorf(format, a...)
	os.Exit(1)
}

func (p *Processor) CaptureEvents(w *fsnotify.Watcher, workers chan struct{}) {

	for {
		select {
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			printTime("ERROR: %s", err)
		case event := <-w.Events:

			if event.Has(fsnotify.Create) || event.Has(fsnotify.Write) || event.Has(fsnotify.Rename) {
				workers <- struct{}{}
				go p.processEvent(event, workers)
			}

		}
	}
}

func (p *Processor) processEvent(event fsnotify.Event, workers <-chan struct{}) {
	defer func() { <-workers }()
	size := int64(0)
	var err error
	log.Info("Processing event: ", event)
	printTime(" %s", event)

	size, err = io.GetFileSize(event.Name)
	if err != nil {
		log.Error(err)
	}

	fmt.Println("Size of the file: ", size)

	writer <- FileEvent{FileName: event.Name, Size: size}

}

func (p *Processor) aggregateForBatch() {
	batchSize := p.config.MaxGoRoutines

	for {
		select {
		case data := <-writer:
			if batch == nil {
				batch = make(map[string]int64)
			}
			batch[data.FileName] = data.Size
			fmt.Print("Batch size: ", len(batch), batchSize)
			if len(batch) >= batchSize {
				filename := "output.json"
				p.appendToJSONFile(batch, filename)
				batch = nil
			}
			if len(batch) > 0 {
				fmt.Print("Batch size: ", len(batch), batchSize)
				filename := "output.json"
				p.appendToJSONFile(batch, filename)
			}

		}
	}

}

func (p *Processor) appendToJSONFile(data map[string]int64, filename string) {
	filePath := filepath.Join(p.config.TargetDirectory, filename)

	content, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to read file: %v", err)
	}
	var existingData map[string]int64
	if len(content) > 0 {
		if err := json.Unmarshal(content, &existingData); err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
	} else {
		existingData = make(map[string]int64)
	}

	for key, value := range data {
		existingData[key] = value
	}

	updatedContent, err := json.Marshal(existingData)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(filePath, updatedContent, 0644); err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

}
