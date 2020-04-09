package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

func test(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("ok\n"))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("[WARNING] 404", r.URL.Path)
	http.Error(w, "404 page not found", http.StatusNotFound)
}

func main() {
	//http.HandleFunc("/", test)
	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		prefix = "/"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	if prefix != "/" {
		http.Handle("/", http.HandlerFunc(NotFound))
	}

	http.Handle(prefix, &webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("WEBDAV [%s]: %s, ERROR: %s\n", r.Method, r.URL, err)
			} else {
				log.Printf("WEBDAV [%s]: %s \n", r.Method, r.URL)
			}
		},
	})
	log.Println("Webdav port:", port, "prefix:", prefix)

	log.Fatal(http.ListenAndServe(port, nil))

}
