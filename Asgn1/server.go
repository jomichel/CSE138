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

		var res string

		s := strings.Split(req.URL.Path[1:], "/")
		if len(s) == 1 {
			w.Header().Set("Content-Type", "application/json")

			res = `{"message": "world"}`
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "Method Not Allowed")
		}

		io.WriteString(w, res)
	case "POST":
		var res string

		s := strings.Split(req.URL.Path[1:], "/")
		if len(s) == 2 {
			w.Header().Set("Content-Type", "application/json")

			res = fmt.Sprintf(`{"message": "Hi, %s."}`, s[1])
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "Method Not Allowed")
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
			http.Error(w, "Internal Server Error", 500)
			return
		}
		if len(rew.Form["msg"]) == 1 {
			s := rew.Form.Get("msg")
			res = fmt.Sprintf(`{"message": "%s"}`, s)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Bad Request")
			return

		}

	}
	io.WriteString(w, res)

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello/", helloHandler)

	http.HandleFunc("/test", testHandler)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
