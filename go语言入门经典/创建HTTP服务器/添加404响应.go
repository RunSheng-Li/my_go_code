package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("跨越长城，走向世界\n"))

}

func main() {

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)

}
