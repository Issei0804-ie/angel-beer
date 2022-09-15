package angel_beer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func NewModel() {
	fmt.Println("model")
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	h3 := func(w http.ResponseWriter, _ *http.Request) {

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/supervisor", SupervisorRequest)
	http.HandleFunc("/agent", h3)
	http.HandleFunc("/agent/ping", PingRequest)
	http.ListenAndServe(":8080", nil)
}

func SupervisorRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		fmt.Println(string(buf.Bytes()))
	}
}

func PingRequest(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Pong!!!!!!!!!!!!!!!!!!!!!!\n")
}
