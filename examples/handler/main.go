package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/houseme/mobiledetect"
)

// Handler .
type Handler struct{}

// Mobile .
func (h *Handler) Mobile(w http.ResponseWriter, r *http.Request, m *mobiledetect.MobileDetect) {
	fmt.Fprint(w, "Hello, this is mobile")
}

// Tablet .
func (h *Handler) Tablet(w http.ResponseWriter, r *http.Request, m *mobiledetect.MobileDetect) {
	fmt.Fprint(w, "Hello, this is tablet")
}

// Desktop .
func (h *Handler) Desktop(w http.ResponseWriter, r *http.Request, m *mobiledetect.MobileDetect) {
	fmt.Fprint(w, "Hello, this is desktop", m.MobileGrade())
}

func main() {
	log.Println("Starting local server http://localhost:10001/check (cmd+click to open from terminal)")
	mux := http.NewServeMux()
	h := &Handler{}
	mux.Handle("/", mobiledetect.Handler(h, nil))
	http.ListenAndServe(":10001", mux)
}
