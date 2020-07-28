package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	port = 8080
)

func main() {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(onRequest),
	}
	log.Printf("Listening on %d", port)

	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func onRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	if r.Method == http.MethodPut {
		echoCapitalHandler(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

type flushWriter struct {
	w io.Writer
}

func (fw flushWriter) Write(p []byte) (n int, err error) {
	log.Printf("Received: %s", p)
	n, err = fw.w.Write(p)
	// Flush - send the buffered written data to the client
	if f, ok := fw.w.(http.Flusher); ok {
		f.Flush()
	}
	return
}

func echoCapitalHandler(w http.ResponseWriter, r *http.Request) {
	// First flash response headers
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	// Copy from the request body to the response writer and flush
	// (send to client)
	io.Copy(flushWriter{w: w}, r.Body)
}
