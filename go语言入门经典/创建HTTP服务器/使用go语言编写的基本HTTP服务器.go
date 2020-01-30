package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("跨越长城，走向世界\n"))
}

func main() {

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)

}

//运行程序后访问 http://localhost:8000/