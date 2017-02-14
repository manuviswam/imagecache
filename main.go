package main

import (
	"net/http"

	h "github.com/manuviswam/imagecache/handler"
	g "github.com/manuviswam/imagecache/gateway"
)

func main() {
	gateway := g.ImageGateway {
	}
	http.HandleFunc("/image/api", h.ImageHandler(&gateway))
	http.ListenAndServe(":8080", nil)
}