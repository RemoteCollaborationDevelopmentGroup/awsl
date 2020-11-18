package main

import (
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("------ http3 server ------")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello quic"))
	})
	go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_host := strings.Split(r.Host, ":")
			_host[1] = "443"
			target := "https://" + strings.Join(_host, ":") + r.URL.Path
			if len(r.URL.RawQuery) > 0 {
				target += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, target, http.StatusTemporaryRedirect)
		}))
	http3.ListenAndServe("0.0.0.0:443", "./ssl/4790070_satori.love.pem", "./ssl/4790070_satori.love.key", nil)
}