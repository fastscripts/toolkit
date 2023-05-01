package main

import (
	"log"
	"net/http"

	"github.com/fastscripts/toolkit"
)

func main() {

	mux := routes()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/download", downloadFile)
	return mux

}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools
	tools.DownloadStaticFile(w, r, "./files", "Bee.png", "bee.png")
}
