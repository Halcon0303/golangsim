package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello my name is\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func serveHTML(w http.ResponseWriter, req *http.Request) {
	htmlContent, err := ioutil.ReadFile("./simple.html")
	if err != nil {
		http.Error(w, "Failed to read HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(htmlContent)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
func serveHTML1(w http.ResponseWriter, req *http.Request) {
	htmlContent, err := ioutil.ReadFile("./simple1.html")
	if err != nil {
		http.Error(w, "Failed to read HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(htmlContent)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/simple", serveHTML)
	http.HandleFunc("/simple1", serveHTML1)

	http.ListenAndServe(":8001", nil)
}
