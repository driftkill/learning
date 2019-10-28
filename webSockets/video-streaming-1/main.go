// !!! Not using WebSockets in this !!!
// Program not working as of now...
// Needs minor changes.

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rangeValue := r.Header.Get("range")
	fmt.Println("Range:")
	fmt.Println(rangeValue)

	buf := bytes.NewBuffer(nil)
	f, _ := os.Open("sample.mp4")
	io.Copy(buf, f)
	f.Close()

	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", "22074728")
	w.Header().Set("Last-Modified", "Mon, 28 Oct 2019")

	w.WriteHeader(206)
	w.Write(buf.Bytes())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7000", nil)
}
