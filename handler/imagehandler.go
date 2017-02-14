package handler

import (
	"net/http"

	p "github.com/manuviswam/imagecache/processor"
)

func ImageHandler(processor *p.ImageProcessor) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename='picture.png'")
		w.Write(processor.GetImageWithSize("https://3.bp.blogspot.com/-W__wiaHUjwI/Vt3Grd8df0I/AAAAAAAAA78/7xqUNj8ujtY/s1600/image02.png", 600, 300))
	}
}