package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("../frontend"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		htmlFilePath := "../frontend/index.html"

		content, err := ioutil.ReadFile(htmlFilePath)
		if err != nil {

			http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "text/html")

		w.Write(content)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, re *http.Request) {
		if re.Method == http.MethodPost {
			fmt.Println("Post request recived")
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("Post request recived and prossed successfully"))
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("Invalid request"))
		}

	})

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
