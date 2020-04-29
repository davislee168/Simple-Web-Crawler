package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf(“GET error: %v”, err)
		return nil, err
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	} else {
       body, err := ioutil.ReadAll(resp.Body)
       if err != nil {
       	log.Printf(“Read error: %v”, err)
       	return nil, err
       } else {
       	return body, nil
       }
	}
}
