package internal

import (
	"file-monitoring-service/internal/config"
	"file-monitoring-service/internal/util/io"
	"fmt"
	"os"

	"time"

	"github.com/fsnotify/fsnotify"
)

type Processor struct {
	config *config.Config
}

type EventProcessor interface {
	UpdateStatus() error
	ProcessEvent() error
	NewProcessor() (*Processor, error)
	ListenFolderEvents() error
}

func NewProcessor() (*Processor, error) {

	path := os.Getenv("CONFIG_PATH")

	fmt.Println("Loading config from ", path)
	config, err := config.LoadConfig(path)

	processor := &Processor{
		config: config,
	}
	if err != nil {
		return nil, err
	}
	return processor, nil
}

func (p *Processor) ProcessEvents() error {

	paths, err := io.GetSubDirectories(p.config.SrcDirectory)
	paths = append(paths, p.config.SrcDirectory)
	if err != nil {
		fmt.Println(err)
		return err

	}

	watch(paths)
	return nil
}

func (p *Processor) UpdateStatus() error {
	//TODO: Implement
	return nil
}

func watch(paths []string) {
	if len(paths) < 1 {
		exit("must specify at least one path to watch")
	}

	// Create a new watcher.
	w, err := fsnotify.NewWatcher()
	if err != nil {
		exit("creating a new watcher: %s", err)
	}
	defer w.Close()

	// Start listening for events.
	go watchLoop(w)

	// Add all paths from the commandline.
	for _, p := range paths {
		err = w.Add(p)
		if err != nil {
			exit("%q: %s", p, err)
		}
	}

	printTime("ready; press ^C to exit")
	<-make(chan struct{}) // Block forever
}

func printTime(s string, args ...interface{}) {
	fmt.Printf(time.Now().Format("15:04:05.0000")+" "+s+"\n", args...)
}

func watchLoop(w *fsnotify.Watcher) {
	i := 0
	for {
		select {
		// Read from Errors.
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}
			printTime("ERROR: %s", err)
		// Read from Events.
		case e, ok := <-w.Events:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}

			// Just print the event nicely aligned, and keep track how many
			// events we've seen.
			i++
			printTime("%3d %s", i, e)
		}
	}
}

func exit(format string, a ...interface{}) {
	os.Exit(1)
}
