package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the main page!"))
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This page is for testing."))
	})

	for i, arg := range os.Args {
		if i != 0 {
			switch arg {
			case "-p":
				if len(os.Args) > i+1 {
					p, _ := strconv.Atoi(os.Args[i+1])
					fmt.Printf("Server is running at 'http://localhost:%d'\n", p)
					http.ListenAndServe(fmt.Sprintf(":%d", p), nil)
				} else {
					fmt.Println("You need to provide an argument to '-d'")
				}
			case "-h":
				fmt.Printf("Usage: %s [-p PORT | -h]\n", os.Args[0])
			}
		}
	}
}
