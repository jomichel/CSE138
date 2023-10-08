package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case "GET":
		// get url path

		var res string

		w.Header().Set("Content-Type", "application/json")
		s := strings.Split(req.URL.Path[1:], "/")
		if len(s) == 1 {
			res = `{"message": "world"}`
		} else {
			http.Error(w, "Method Not Allowed", 405)

		}

		io.WriteString(w, res)
	case "POST":
		var res string

		w.Header().Set("Content-Type", "application/json")
		s := strings.Split(req.URL.Path[1:], "/")
		if len(s) == 2 {
			res = fmt.Sprintf(`{"message": "Hello, %s"}`, s[1])
		} else {
			http.Error(w, "Method Not Allowed", 405)
		}
		io.WriteString(w, res)

	}

}
func testHandler(w http.ResponseWriter, rew *http.Request) {
	var res string
	switch rew.Method {
	case "GET":
		res = fmt.Sprintf(`{"message": "%s"}`, "test is successful")

	case "POST":
		if err := rew.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			http.Error(w, "Internal Server Error", 500)

			return
		}
		if len(rew.Form["msg"]) == 1 {
			s := rew.Form.Get("msg")
			res = fmt.Sprintf(`{"message": "%s"}`, s)
		} else {
			http.Error(w, "Bad Request", 400)
			return

		}

	}
	io.WriteString(w, res)

}

func main() {
	// Hello world, the web server

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello/", helloHandler)

	http.HandleFunc("/test", testHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
