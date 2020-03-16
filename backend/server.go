package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func envDefault(key, or string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return or
}

var (
	confSavePath = envDefault("TA_SAVE_PATH", "./save")
	confAssets = envDefault("TA_ASSETS", "../frontend/public")
	confListenAddr = envDefault("TA_LISTEN_ADDR", ":8484")
)

func main() {
	// serve json updating
	http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		out, err := os.Create(confSavePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("creating save path: %v", err), 500)
			return
		}
		if _, err := io.Copy(out, r.Body); err != nil {
			http.Error(w, fmt.Sprintf("writing save: %v", err), 500)
			return
		}
	})
	// serve json fetching
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, confSavePath)
	})
	// serve frontend
	http.Handle("/", http.FileServer(http.Dir(confAssets)))
	//
	log.Printf("starting on %q\n", confListenAddr)
	if err := http.ListenAndServe(confListenAddr, nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
