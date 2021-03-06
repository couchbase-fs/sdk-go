package main

import (
	"log"
	"net/http"
)

func doExport(w http.ResponseWriter, req *http.Request,
	path string) {

	quit := make(chan bool)
	defer close(quit)
	ch := make(chan *namedFile)
	cherr := make(chan error)

	go pathGenerator(path, ch, cherr, quit)
	go logErrors("export", cherr)

	err := streamFileMeta(w, ch, cherr)
	if err != nil {
		log.Printf("Error exporting meta: %v", err)
	}
}
