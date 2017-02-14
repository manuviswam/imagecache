package main

import (
	"net/http"

	h "github.com/manuviswam/imagecache/handler"
	g "github.com/manuviswam/imagecache/gateway"
	c "github.com/manuviswam/imagecache/cache"
	p "github.com/manuviswam/imagecache/processor"
)

func main() {
	gateway := &g.ImageGateway {
	}
	cache := c.InitCache()
	processor := p.ImageProcessor{
		ImageCache : cache,
		Gateway : gateway,
	}
	http.HandleFunc("/image/api", h.ImageHandler(&processor))
	http.ListenAndServe(":8080", nil)
}