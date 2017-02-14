package handler

import (
	"net/http"

	g "github.com/manuviswam/imagecache/gateway"
)

func ImageHandler(gateway *g.ImageGateway) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename='picture.png'")
		w.Write(gateway.Get("https://3.bp.blogspot.com/-W__wiaHUjwI/Vt3Grd8df0I/AAAAAAAAA78/7xqUNj8ujtY/s1600/image02.png"))
	}
}