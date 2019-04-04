package main

import (
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("ERROR", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			// watch for events
			case _ = <-watcher.Events:
				log.Println("INFO", "change event")

				db, err := sql.Open("postgres", os.Getenv("PGBOUNCER_URI"))
				if err != nil {
					log.Println("ERROR", err)
					continue
				}

				_, err = db.Exec("RELOAD")
				if err != nil {
					log.Println("ERROR", err)
				}

				db.Close()

			// watch for errors
			case err := <-watcher.Errors:
				log.Println("ERROR", err)
			}
		}
	}()

	if os.Getenv("PGBOUNCER_CONFIG") == "" {
		log.Println("ERROR", "missing variable PGBOUNCER_CONFIG")
		close(done)
	} else if os.Getenv("PGBOUNCER_URI") == "" {
		log.Println("ERROR", "missing variable PGBOUNCER_URI")
		close(done)
	}

	for _, file_path := range strings.Split(os.Getenv("PGBOUNCER_CONFIG"), ";") {
		if err := watcher.Add(file_path); err != nil {
			log.Println("ERROR", err)
			close(done)
			break
		} else {
			log.Println("INFO", fmt.Sprintf("Watching '%s'", file_path))
		}
	}

	<-done
}
