package gateway

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

type ImageGateway struct {
}

func (g *ImageGateway) Get(url string) []byte {
	fmt.Println("Making request to ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching image ", err)
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response ", err)
		return nil		
	}

	return body
}