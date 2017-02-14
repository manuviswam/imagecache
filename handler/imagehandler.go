package handler

import (
	"net/http"
	"strconv"

	p "github.com/manuviswam/imagecache/processor"
)

func ImageHandler(processor *p.ImageProcessor) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		width,_ := strconv.Atoi(r.URL.Query().Get("width"))
		height,_ := strconv.Atoi(r.URL.Query().Get("height"))
		w.Header().Set("Content-Disposition", "attachment; filename='picture.png'")
		w.Write(processor.GetImageWithSize(url, width, height))
	}
}