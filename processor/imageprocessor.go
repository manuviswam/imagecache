package processor

import (
	"crypto/sha1"
	"strconv"
	"image/png"
	"bytes"
	"fmt"

	g "github.com/manuviswam/imagecache/gateway"
	c "github.com/manuviswam/imagecache/cache"
	r "github.com/nfnt/resize"
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
		p.ImageCache.Add(key, img)
	}

	resizedImg := resize(img, width, height)
	p.ImageCache.Add(keyWithSize, resizedImg)
	
	return resizedImg
}

func resize(img []byte, width, height int) []byte {
	reader := bytes.NewReader(img)
	decodedImage, err := png.Decode(reader)
	if err != nil {
		fmt.Println("Error decoding image ", err)
		return nil
	}

	resizedImage := r.Resize(uint(width), uint(height), decodedImage, r.Lanczos3)
	outBuffer := new(bytes.Buffer)
	png.Encode(outBuffer, resizedImage)
	return outBuffer.Bytes()
}