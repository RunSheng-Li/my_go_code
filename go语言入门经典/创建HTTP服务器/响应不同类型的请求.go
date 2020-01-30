package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("跨越长城，走向世界\n"))

	switch r.Method {

	case "GET":
		w.Write([]byte("收到了一个GET请求\n"))

	case "POST":
		w.Write([]byte("收到了一个POST请求\n"))

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
	}

}

func main() {

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8001", nil)

}
