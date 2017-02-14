package processor

import (
	"crypto/sha1"
	"strconv"

	g "github.com/manuviswam/imagecache/gateway"
	c "github.com/manuviswam/imagecache/cache"
)

type ImageProcessor struct {
	ImageCache *c.Cache
	Gateway *g.ImageGateway
}

func (p *ImageProcessor) GetImageWithSize(url string, width, height int) []byte {
	h := sha1.New()
	h.Write([]byte(url))

	key := string(h.Sum(nil))
	keyWithSize := key + strconv.Itoa(width) + "x" + strconv.Itoa(height)

	if p.ImageCache.IsPresent(keyWithSize) {
		return p.ImageCache.Get(keyWithSize)
	}
	var img []byte

	if p.ImageCache.IsPresent(key) {
		img = p.ImageCache.Get(key)
	} else {
		img = p.Gateway.Get(url)
	}

	resizedImg := resize(img, width, height)
	p.ImageCache.Add(keyWithSize, resizedImg)
	
	return resizedImg
}

func resize(img []byte, widh, height int) []byte {
	return img
}