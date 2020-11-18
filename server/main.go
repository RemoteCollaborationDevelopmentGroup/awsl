package main

import (
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("------ http3 server ------")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello quic"))
	})
	http3.ListenAndServe("0.0.0.0:443", "./ssl/4790070_satori.love.pem", "./ssl/4790070_satori.love.key", nil)
}