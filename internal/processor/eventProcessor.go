package internal

import (
	"file-monitoring-service/internal/config"
	"file-monitoring-service/internal/util/io"
	"fmt"
	"os"

	"time"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

type Processor struct {
	config *config.Config
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

func (p *Processor)watch(paths []string) {
	if len(paths) < 1 {
		exit("must specify at least one path to watch")
	}
	w, err := fsnotify.NewWatcher()
	events := make(chan fsnotify.Event)
	go p.CaptureEvents(events)
	if err != nil {
		exit("creating a new watcher: %s", err)
	}
	defer w.Close()
	go watchLoop(w, events)

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

func watchLoop(w *fsnotify.Watcher, events chan<- fsnotify.Event) {
	i := 0
	for {
		select {
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			printTime("ERROR: %s", err)

		case event, ok := <-w.Events:
			if !ok {
				return
			}
			i++

			events <- event

		}
	}
}

func exit(format string, a ...interface{}) {
	log.Errorf(format, a...)
	os.Exit(1)
}

func (p *Processor) CaptureEvents(events <-chan fsnotify.Event) error {

	workers := make(chan struct{}, p.config.MaxGoRoutines)

	for {
		select {
		case event := <-events:
			workers <- struct{}{}
			go p.processEvent(event, workers)
		}
	}
}

func (p *Processor) processEvent(event fsnotify.Event, workers <-chan struct{}) {
	defer func() { <-workers }()
	log.Info("Processing event: ", event)
	printTime(" %s", event)
}
